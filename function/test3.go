package function

import (
	"fmt"
)

func Test3() {
	a, b := Split(9)
	fmt.Println(a, b)
	var c string = "test"
	fmt.Println(&c)
}

/*
	func swap(x, y string) (string, string) {
		return y, x
	}
*/
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func GetValue(pointer *int) int {
	return *pointer
}
