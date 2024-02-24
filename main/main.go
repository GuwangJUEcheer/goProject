package main

//这个包一定要是main包
//编译采用命令 go build -o bin xxx.exe(这里自定义) example.com/myproject/main(main所在路径前面gopath)
import (
	learnSlice "example.com/myproject/stringdemo"
)

func main() {
	learnSlice.ShowSlice()
	learnSlice.Test()
}
