package function

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

type Computer struct {
}

//实现
func (p Phone) Start() {

	fmt.Println("手机开始工作")
}

func (c Computer) Start() {

	fmt.Println("电脑开始工作")
}

func (c Computer) Working(usb Usb) {

	usb.Start()
	usb.Stop()
}

//高内聚低耦合
type StudentInterface interface {
	Say()
}

type StudentChin struct {
}

func (stud StudentChin) Say() {

	fmt.Println("student say")
}
func TestInterface() {

	stud := StudentChin{} //结构体变量实现了StudentInterface接口

	var a StudentInterface = stud

	a.Say()
}

//只要是自定义变量都可以实现接口 接口里面所有的方法都必须实现 只有实现了接口所有的方法才能这么写 such as :TestInterface
type StudentInterface2 interface {
	StudentInterface
	Test()
}

//实现上面接口的话就要实现StudentInterface 里面所有方法 和 Test
