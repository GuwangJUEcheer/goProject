package main

//这个包一定要是main包
//编译采用命令 go build -o bin xxx.exe(这里自定义) example.com/myproject/main(main所在路径前面gopath)
//引入的最先
import (
	"fmt"

	learnSlice "example.com/myproject/stringdemo"

	"example.com/myproject/utils"

	"strings"
)

//累加器

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
	//匿名函数 没有名字的函数就是匿名函数 一般只使用一次
	res1 := func(a1, a2 int) int {
		return a1 + a2
	}(15, 25)
	fmt.Printf("res1 = %d\n", res1)
	fmt.Printf("res2 = %d\n", utils.Fun1(3, 9))
	f := add()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //12

	f1 := makeSuffix(".jpg")

	fmt.Println(f1("00test"))
	sum(1, 2)
	utils.Test()
}

// 定义一个闭包
func add() func(int) int {
	var n int = 10
	return func(n1 int) int {

		return n + n1
	}
}

// 闭包的一个小例子 里面局部变量只初始化一次
func makeSuffix(suffix string) func(name string) string {

	return func(name string) string {

		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return suffix
	}
}

// defer关键字 有点像.net的await关键字 最关键 函数执行完之后自动释放资源
/* defer file.close() defer connect.close*/
func sum(n1, n2 int) int {

	defer fmt.Println(n1) //按照先进后出 所以最后执行
	defer fmt.Println(n2)
	fmt.Println(n1 + n2) //最早执行
	return n1 + n2
}
