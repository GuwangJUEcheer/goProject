package function

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	name string
	sex  string
	age  int
}

type StudentO struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int32  `json:"age"`
}

// 小写字母开头的结构体实例化使用工厂模式
type person struct {
	name string
	sex  string
	age  int
}

func NewPerson(name string, sex string, age int) *person {

	return &person{
		name,
		sex,
		age,
	}
}

func (p *person) GetName() string {

	return p.name
}

func TestStruct() {
	var stud Student
	stud.name = "tom"
	stud2 := stud
	stud2.name = "jack"
	fmt.Println(stud2.name) //互不干扰 它是值类型的只有引用类型才会跟着变

	var stud3 Student = Student{"hary", "male", 20}
	fmt.Printf("sui = %v  \t", stud3.age)

	//指针赋值
	var stud4 *Student = new(Student)
	(*stud4).name = "jackson" //也可以直接stud4.name
	fmt.Printf("name = %v \t", stud4.name)

	var stud5 *Student = &Student{}
	stud5.name = "terry"
	fmt.Printf("name = %v \t", stud5.name)

	/*
		结构体在内存中是连续的
		可以type重新声明类型 可以强制转换
	*/
	type stu Student
	var stud6 stu = stu(stud3)
	jsonstring, err := json.Marshal(stud6) //这里因为不是大写所以导致错误了
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Println(jsonstring)
	}

	var studO StudentO = StudentO{
		"tom",
		"sasa",
		12,
	}

	jsonstr, err := json.Marshal(studO) //这里因为不是大写所以导致错误了
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonstr)) //这里要转换的
	}

	stud3.TestStudent() //给结构体绑定方法
	// 方法和函数区别 当参数是函数指针时候函数必须传地址 但是方法可以传类型
	//比如上面的 TestStudent
	(&stud3).TestStudent() //两个方法执行结果完全一致 因为执行是值拷贝

	var stud7 = &Student{"aa", "male", 20}
	fmt.Printf("%v", stud7)

}
func (s Student) TestStudent() {

	println("test")
}
