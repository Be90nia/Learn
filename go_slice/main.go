package main

import "fmt"

//切片
func main() {
	var slice []int
	fmt.Printf("len:%d, cap:%d\n", len(slice), cap(slice))
}
