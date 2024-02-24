package function

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func rangeArray() {

	// 定义一个 map，键为字符串，值为整数
	oldMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// 创建一个新的空 map，用于存储 oldMap 的复制
	newMap := make(map[string]int)
	fmt.Println(pow)
	// 将 oldMap 的内容复制到 newMap 中
	for key, value := range oldMap {
		newMap[key] = value
	}

	// 打印新的 map，验证复制是否成功
	fmt.Println(newMap)

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	/*
	   for key := range oldMap
	   for key := range oldMap
	*/
}
