package tool

import (
	"bufio"
	"fmt"
	"generateStruct/tool"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var (
	packageName          = "package %s\n\n"
	structName           = "type %s struct {"
	structValue          = "    %s %s `json:\"%s\"`"
	structValueWithNotes = "    %s %s `json:\"%s\"`  // %s"
	structEnd            = "}\n"
	wg                   sync.WaitGroup
)

type MakeMsg struct {
	allReadPath []string
}

// 读取msg文件
func (this *MakeMsg) ReadMsg(readPath, writePath string) error {
	if readPath == "" || writePath == "" {
		return fmt.Errorf("ReadMsg|readPath is nil")
	}
	files, err := ioutil.ReadDir(readPath)
	if err != nil {
		return fmt.Errorf("ReadMsg|readDir is err:%v", err)
	}
	this.allReadPath = make([]string, 0)
	allWritePath := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		secondPath := readPath + "\\" + file.Name()
		secondFiles, err := ioutil.ReadDir(secondPath)
		if err != nil {
			return fmt.Errorf("ReadMsg|secondFiles:%v readDir is err:%v", secondPath, err)
		}
		for _, secondFile := range secondFiles {
			thirdPath := secondPath + "\\" + secondFile.Name()
			if path.Ext(thirdPath) != ".txt" {
				continue
			}
			this.allReadPath = append(this.allReadPath, secondPath+"\\"+secondFile.Name())
			allWritePath = append(allWritePath, writePath+"\\"+file.Name())
		}
	}

	wg.Add(len(this.allReadPath))
	allWritePathLen := len(allWritePath)
	for k, file := range this.allReadPath {
		go func(k int, file string) {
			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("ReadMsg|os.Open is err:%v", err)
			}
			var content string
			read := bufio.NewReader(f)
			for {
				line, _, err := read.ReadLine()
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Printf("ReadMsg|ReadLine is err:%v", err)
					return
				}
				str, err := this.SplicingData(string(line))
				if err != nil {
					continue
				}
				content += string(str) + "\n"
			}
			if k >= allWritePathLen {
				fmt.Printf("err:ReadMsg|k is %d >= allWritePathLen:%d", k, allWritePathLen)
				return
			}
			err = this.WriteMsg(content, file, writePath, allWritePath[k])
			if err != nil {
				fmt.Printf("err:%v\n", err)
			}
			wg.Done()
		}(k, file)
	}
	wg.Wait()
	return nil
}

// 拼装struct
func (this *MakeMsg) SplicingData(value string) (string, error) {
	if value == "" || strings.TrimSpace(value) == "" {
		return "", fmt.Errorf("SplicingData|value is nil")
	}
	if tool.JudgeIndex(value, "{") {
		slice := strings.Split(value, "{")
		if len(slice) != 2 {
			return "", fmt.Errorf("SplicingData|slice[1]:%v len is:%d not 2", slice, len(slice))
		}
		return fmt.Sprintf(structName, strings.TrimSpace(slice[0])), nil
	}
	if tool.JudgeIndex(value, "}") {
		return structEnd, nil
	}
	slice := strings.Split(value, ":")
	if len(slice) != 2 {
		// 判断是否存在结构体上方的注释
		return value, nil
	}
	sliceOne := strings.TrimSpace(slice[0])
	sliceTwo := strings.TrimSpace(slice[1])
	sliceNotes := strings.Split(slice[1], "//")
	if len(sliceNotes) != 2 {
		return fmt.Sprintf(structValue, tool.FirstRuneToUpper(sliceOne), sliceTwo, sliceOne), nil
	}
	return fmt.Sprintf(structValueWithNotes, tool.FirstRuneToUpper(sliceOne), strings.TrimSpace(sliceNotes[0]), sliceOne, strings.TrimSpace(sliceNotes[1])), nil
}

// 生成msg文件
func (this *MakeMsg) WriteMsg(content, readPath, baseWritePath, completeWritePath string) error {
	if content == "" {
		return fmt.Errorf("WriteMsg|content is nil")
	}
	// 获取文件名
	_, fileName := filepath.Split(readPath)
	slice := strings.Split(fileName, ".txt")
	if len(slice) != 2 {
		return fmt.Errorf("WriteMsg|slice's len is %d not 2", len(slice))
	}
	var data string
	data += fmt.Sprintf(packageName, filepath.Base(baseWritePath))
	data += content
	// 判断文件夹是否存在 如果不存在 就新建 如果存在 则不处理
	err := os.MkdirAll(completeWritePath, 0644)
	if err != nil {
		return fmt.Errorf("WriteMsg|os.Mkdir is err:%v", err)
	}
	fw, err := os.OpenFile(completeWritePath+"\\"+slice[0]+".go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("WriteMsg|OpenFile is err:%v", err)
	}
	defer fw.Close()
	_, err = fw.Write([]byte(data))
	if err != nil {
		return fmt.Errorf("WriteMsg|Write is err:%v", err)
	}
	return nil
}
