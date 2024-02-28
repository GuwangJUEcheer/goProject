package main

import (
	"fmt"

	function "example.com/myproject/function"
)

func main() {
	for i, v := range function.Arr3 {
		for j, v := range v {
			fmt.Printf("arr3[%v][%v]=%v \t", i, j, v)
		}
	}
	function.DefineMap()
	function.SortMap()
	function.TestStruct()
}
