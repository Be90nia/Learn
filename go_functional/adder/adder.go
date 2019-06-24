package main

import "fmt"

//闭包
// func adder() func(int) int {
// 	sum := 0                 //局部变量(相对于下面的函数体为自由变量)
// 	return func(v int) int { //闭包
// 		sum += v //自由变量
//      通过自由变量进行不断的自动关联的结构,最终完全将其返回的方式叫闭包(可以查看闭包闭包概念图)
// 		return sum
// 	}
// }

//函数体
func adder() func(int) int {
	sum := 0                 //局部变量
	return func(v int) int { //闭包
		sum += v //自由变量
		return sum
	}
}

//较为正统的函数编程方式
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d = %d\n", i, a(i))
	}

	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0+1+...+%d = %d\n", i, s)
	}
}
