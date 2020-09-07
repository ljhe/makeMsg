package main

import (
	"flag"
	"fmt"
	"makeMsg/tool"
)

var (
	readPath  = flag.String("readPath", "", "the path for read .tet file")
	writePath = flag.String("writePath", "", "the path for write .go file")
)

func main() {
	flag.Parse()
	if *readPath == "" || *writePath == "" {
		fmt.Printf("err:readPath or writePath is nil\n")
		return
	}
	msg := tool.MakeMsg{}
	err := msg.ReadMsg(*readPath, *writePath)
	if err != nil {
		fmt.Printf("something err:%v\n", err)
		return
	}
}
