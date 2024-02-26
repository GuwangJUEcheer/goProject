package utils

/*
   字符串的一些常用函数
*/
import (
	"fmt"
	"strconv"
	"strings"
)

//

func Test() {
	//1 length

	var a string = "test string 北京"

	fmt.Println(len(a))

	b := []rune(a) //避免中文不显示 rune的切片

	for i := 0; i < len(b); i++ {
		fmt.Printf("%c", b[i])
	}

	fmt.Println(toInteger("12"))
	fmt.Printf("%T\n", toString(3))
	fmt.Printf("%v", toByteArray("aaaan"))

	//检查包含某个字串
	fmt.Printf("%v", strings.Contains("aaaasas", "asa"))

	//检查某个字符串有几个字串
	fmt.Printf("%v", strings.Count("nusadaodsasasasas", "a"))

	//不区分大小写字符串比较
	fmt.Printf("%v\n", strings.EqualFold("abc", "ABC"))

	//获得第一次出现Index or -1 获得最后一个用LastIndex
	fmt.Printf("%v\n", strings.Index("abc_guwangje", "_"))

	//替换某个字串(1 代表替换几个)
	fmt.Printf("%v\n", strings.Replace("abc_guwangje", "_", " dot ", 1))

	//大写小写转换 大写就是ToUppper
	fmt.Printf("%v\n", strings.ToLower("ABC"))

	//拆分字串
	fmt.Println(strings.Split("Hello,tani,woaini", ","))

	//去除首尾空格 Trim去除指定字符
	fmt.Println(strings.TrimSpace(" a "))

	//以什么开头 或者以什么结尾
	fmt.Printf("%v ", strings.HasPrefix("guwangjue", "gu"))
	fmt.Printf("%v ", strings.HasSuffix("guwangjue", "jue"))
}

// 字符串转整数
func toInteger(str string) int {

	n, err := strconv.Atoi(str)
	if err != nil {

		fmt.Println("转换错误了", err)
		return 0
	} else {
		fmt.Println(n)
		return n
	}
}

// 整数转字符串
func toString(n int) string {

	return strconv.Itoa(n)
}

// 转换成二进制切片
func toByteArray(str string) []byte {
	return []byte(str)
}

//从二进制转成string

func BytesToString(bytes []byte) string {

	return string(bytes)
}

//formatInt(2) 转为二进制
