@echo off

::将.txt文件转化为相应的.go文件

::读取.txt文件夹路径 会读取该目录下所有文件夹下的.txt文件
set readPath=F:\code\src\makeMsg\test
::生成.go文件的路径 会根据上方读取路径将文件保存到相应文件夹下
set writePath=F:\code\src\makeMsg\tool
::生成可执行文件的路径
set exePath=F:\code\src\makeMsg

echo 开始生成.go文件
%exePath%\makeMsg -writePath=%writePath% -readPath=%readPath%

echo 生成完毕，按任意键继续
TIMEOUT /T 999