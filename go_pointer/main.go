package main

import "fmt"

//go语言的所有函数都是值传递,java、Python都是引用传递
//go语言的指针是不能做运算
func main() {
	a, b := 3, 4
	swap1(a, b)
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)
}

func swap1(a, b int) {
	b, a = a, b
}

func swap2(a, b *int) {
	*b, *a = *a, *b
}
