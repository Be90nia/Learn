package main

import "fmt"

//函数体
func adder() func(int) int {
	sum := 0                 //局部变量
	return func(v int) int { //闭包
		sum += v //自由变量
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d = %d\n", i, a(i))
	}
}
