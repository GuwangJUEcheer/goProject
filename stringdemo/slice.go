package learnSlice

import "fmt"

//必须大写
func ShowSlice() {

	/*
		不只是有关于切片 还有关于make函数的用法
	*/
	//声明切片 有点类似于arraylist
	var slice []int = make([]int, 5, 19) //长度为5 容量为19
	fmt.Println(len(slice))
	//定义一个数组 并且演示数组到切片
	arr := [3]int{1, 2, 3}

	var number = make([]int, 0, 10)
	//切片append之后超过指定的capacity也没有关系golang底层会追加的
	//len →长度 cap →capacity`
	s := []int{1, 2, 3}
	fmt.Println(len(s)) // 输出: 3
	fmt.Println(cap(s)) // 输出: 3
	copy(s, number)
	s = append(s, 4, 5, 6)
	fmt.Println(len(s)) // 输出: 6
	fmt.Println(cap(s)) // 输出: 6

	s = append(s, 7, 8, 9, 10, 11)
	fmt.Println(len(s)) // 输出: 11
	fmt.Println(cap(s)) // 输出: 12
	s = append(s, s...) //复制一份给s
	fmt.Println(cap(s)) // 输出: 24
	//nil的切片
	var numbers []int //default nil
	fmt.Println(numbers)
	//substring
	//numbers :=numbers[:3] //just like arr　这里是左闭右开

	//copy s3 to number
	slice2 := arr[:]
	fmt.Println(slice2)

	/*
		     make定义的不可见 数组直接定义可见
			 slice是引用类型 引用数据发生变化 slice也会变化
			 string和slice一样底层都是数组也可以处理
	*/
	fmt.Println("导包成功了")

	var str string = "导包成功了"
	strSlice := str[:] //不能strSlice[0]= xx这样修改 string不可变
	fmt.Println(strSlice)
}
