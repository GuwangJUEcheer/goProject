package function

import (
	"fmt"
)

func arryaDisplay() {
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
