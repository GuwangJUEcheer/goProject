package utils

/*
   日期的一些常用函数
*/
import (
	"fmt"
	"time"
)

var now time.Time = time.Now()

func TestTime() {

	fmt.Println(time.Now().Weekday())

	//格式化时间
	datestr := fmt.Sprintf("当前年月日 %d-%d-%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Printf("time = %v　　", datestr)
	//第二种写法 必须2006/01/02 15:04:05的任一部分
	fmt.Print(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Print(now.Format("2006/01/02"))
	fmt.Println()
	fmt.Print(now.Format("15:04:05"))
}

// 休眠时候使用时间常量 time (NanoSecond MicroSecond Millisecond Second Minute Hour)
func TimeSleep() {
	fmt.Println("--------start---------")
	fmt.Print(now.Format("2006/01/02 15:04:05"))
	time.Sleep(2 * time.Second)
	fmt.Println("")
	fmt.Println("--------end---------")
	fmt.Print(now.Format("2006/01/02 15:04:05"))
}

// 随机秒Unix UnixNano
func UseUnix() {
	fmt.Printf("unix result=%v unixNano result=%v", now.Unix(), now.UnixMicro())
}
