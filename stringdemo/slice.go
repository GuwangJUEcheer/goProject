package learnSlice
/*
不只是有关于切片 还有关于make函数的用法
*/
//声明切片 有点类似于arraylist
var slice []type = make([]int,5,19) //长度为5 容量为19

m := make(map[string] int) //字符串到整数的映射

ch := make(chan int) //建立int类型的通道

s0 := [] int {1,2,3}

//定义一个数组 并且演示数组到切片
arr := [3]int{1,2,3}
s1:= arr[:]

//截取数组一部分使之成为切片
s2 := arr[1:2] 
s3 := arr[1:]
s4 := arr[:3]

//切片append之后超过指定的capacity也没有关系golang底层会追加的
//len →长度 cap →capacity`
s := []int{1, 2, 3}
fmt.Println(len(s)) // 输出: 3
fmt.Println(cap(s)) // 输出: 3

s = append(s, 4, 5, 6)
fmt.Println(len(s)) // 输出: 6
fmt.Println(cap(s)) // 输出: 6

s = append(s, 7, 8, 9, 10, 11)
fmt.Println(len(s)) // 输出: 11
fmt.Println(cap(s)) // 输出: 12

//nil的切片
var numbers []int //default nil

//substring
numbers[:3] //just like arr

//copy s3 to number
copy(s3,number)

