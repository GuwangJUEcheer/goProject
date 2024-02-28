package function

import (
	"fmt"
	"sort"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func RangeArray() {

	// 定义一个 map，键为字符串，值为整数 map key可以是chan也可以是指针但不可以是function
	//map一定要make分配内存后才可以使用
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

func DefineMap() {
	a := make(map[string]string, 10) //10对
	a["name1"] = "jack"
	a["name2"] = "time"
	a["name1"] = "jackson" //会重复
	fmt.Println(a)

	cites := make(map[string]string)
	cites["wuhan"] = "武汉"
	cites["bejing"] = "北京"
	cites["xingtai"] = "邢台"
	fmt.Println(cites)

	person := map[string]string{
		"no1": "aa",
		"no2": "bb",
	}
	fmt.Println(person)

	//当value是map时候
	studentMap := make(map[string]map[string]string) //一定要有make
	studentMap["stud01"] = person
	studentMap["stud02"] = person
	studentMap["stud01"]["no1"] = "tani"
	fmt.Println(studentMap)

	//删除map 不存在key也不会报错
	delete(studentMap, "stud02")
	fmt.Println(studentMap)

	//map查找
	res, ok := cites["wuhan"]
	fmt.Printf("res = %v,ok = %v", res, ok)

	//切片和map一起使用
	sliceMap := make([]map[string]string, 2)
	if sliceMap[0] == nil {
		sliceMap[0] = make(map[string]string, 2)
		sliceMap[0]["name"] = "yutu"
		sliceMap[0]["age"] = "jo"
	}

	//使用append可以动态增加
	//先定义新对象
	newMap := map[string]string{
		"age":  "10",
		"name": "sasa",
	}
	sliceMap = append(sliceMap, newMap)
	fmt.Println(sliceMap)
}

func SortMap() {
	//map是无序的
	var mapsorted map[int]string = map[int]string{
		1:  "aaa",
		10: "bbb",
		3:  "vccc",
		2:  "ddd",
	}

	var keys []int
	for k, _ := range mapsorted {
		keys = append(keys, k)
	}
	sort.Ints(keys) //排序

	//遍历按照key输出值
	for _, v := range keys {
		fmt.Printf("value = %v \t", mapsorted[v])
	}
}
