package main

import (
	"fmt"
	"unsafe"
)

func main() {

	/*
		关于char
	*/
	var c1 byte = 'a'
	var c2 byte = 'b'
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	fmt.Println("c2=", unsafe.Sizeof(c1))

	/*
		关于String
	*/

	var str string = "Hello,world"
	fmt.Println(str)

	var appendString string = "hello" + "world" +
		"cherry"
	fmt.Println(appendString)

	/*
		基本数据类型转换
	*/

	//1 必须所有的都要显示转换
	i := 100
	n1 := float32(i)
	fmt.Printf("n1=%v", n1)
}
