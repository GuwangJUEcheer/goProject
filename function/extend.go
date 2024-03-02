package function

import "fmt"

type Goods struct {
	Name  string
	Price float64
}

// Book可以使用 所有的
type Book struct {
	Goods //嵌套结构体

	Write string
}

type Writer struct {
	Books  *Book //有名结构体(如果有名字必须得带上名字)
	Writer string
}

func TestExtend() {

	//也可以这么写 更快访问
	testObj := Writer{

		&Book{
			Goods{
				"aab",
				12,
			},
			"ffe",
		},
		"cdd",
	}

	fmt.Println(testObj.Books.Name)
}
