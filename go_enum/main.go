package main

import "fmt"

//创建一个新的类型
type Sate int

//使用const以及iota自动赋值
const (
	Running    Sate = iota // value --> 0
	Stopped                // value --> 1
	Rebooting              // value --> 2
	Terminated             // value --> 3
)

//Sate 类型就能当做枚举使用了
func (this Sate) String() string {
	switch this {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	default:
		return "Unknow"
	}
}

func main() {
	var state Sate
	fmt.Println("state", state)
	state = Stopped
	fmt.Println("state", state)
}
