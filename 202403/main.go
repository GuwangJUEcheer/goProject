package main

import (
	"fmt"

	function "example.com/myproject/function"
)

func main() {
	var b function.Book = function.Book{
		Goods: function.Goods{
			Name:  "aa",
			Price: 15.5,
		},
		Write: "tom",
	}
	fmt.Println(b.Name) // 这里应该访问b.Goods.Name，因为Name是Goods结构体的字段

	function.TestExtend()
	function.TestInterface()
}
