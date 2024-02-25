package utils

//演示创建init

var Age int
var Name string

//全局匿名函数
var (
	// fun1是全局匿名函数 要大些
	Fun1 = func(n1, n2 int) int {
		return n1 * n2
	}
)

func init() {
	Age = 1000
	Name = "san"
}
