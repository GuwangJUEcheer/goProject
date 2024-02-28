package function

import (
	"fmt"
)

func ArryaDisplay() {
	var arr [5]int

	for i := 0; i < len(arr); i++ {
		arr[i] = 2 * i
	}
	fmt.Print(arr)

	for i, v := range arr {
		fmt.Printf("索引 %d 的值为 %d\n", i, v)
	}

	arr[2] = 100
	fmt.Println("修改后的数组:", arr)

	anotherOne := [4]string{"i", "love", "yiyi", "!"}
	fmt.Println("另一个数组", anotherOne)

	slice := anotherOne[1:4]
	fmt.Println("追加元素后的切片:", slice)

	// 通过切片修改原始数组
	slice[0] = "tani"
	fmt.Println("修改后的原始数组:", anotherOne)
}

func GetArrayAddress() {

	var intArr [3]int

	// %p可以打印地址 数组地址直接指针即可 数组元素地址最快 int占8 而且连续
	fmt.Printf("%p,%p,%p", &intArr[0], &intArr[1], &intArr[2])
}

func PrintArr() {

	num3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(num3)

	num4 := [...]string{2: "tom", 0: "jack", 1: "Mick"}
	for i, v := range num4 {

		fmt.Printf("i=%v , value=%v ", i, v)
	}
}

/*DoubleArray 二维数组*/
var Arr [4][6]int
var Arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
