package main

//这个包一定要是main包
//编译采用命令 go build -o bin xxx.exe(这里自定义) example.com/myproject/main(main所在路径前面gopath)
//引入的最先
import (
	"fmt"

	learnSlice "example.com/myproject/stringdemo"

	"example.com/myproject/utils"
)

var meth = test()

func test() int {

	fmt.Println("全局变量定义")
	return 0
}

// 全局变量定义最先 其次init
func init() {

	fmt.Println("这个函数在main前调用")
	fmt.Println(meth)
	//类似于string.format
	fmt.Printf("age = %d 岁\n", utils.Age)
}

func main() {
	learnSlice.ShowSlice()
	learnSlice.Test()
}
