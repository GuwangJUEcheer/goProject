package function

import (
	"errors"
	"fmt"
)

/*
defer panic recover
*/
func TestError() {

	defer func() {
		error := recover() //可以捕获异常
		if error != nil {
			fmt.Println("error =", error)

		}
	}()
	num1 := 10
	num2 := 0

	fmt.Println(num1 / num2)
}

// 自定义异常
func NewError() {

	err := errors.New("error happens")

	panic(err)

}
