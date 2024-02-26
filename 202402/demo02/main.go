package main

import (
	"fmt"

	function "example.com/myproject/function"
)

// 一些内置函数
func main() {

	fmt.Println(len("aasdada"))

	//new关键字
	num := new(int) //类型指针 值指针地址
	fmt.Printf("Num2 type %T num2 value %v num2 true value %v", num, num, *num)
	//function.GetArrayAddress()
	//function.TestError()
	//function.NewError()
	function.PrintArr()

}
