package main

import "fmt"

//切片
func main() {
	var slice []int
	fmt.Printf("len:%d, cap:%d\n", len(slice), cap(slice))

	//切片基础语法
	//切片初始化
	//切片名字 = 数组名[开始位置:结束位置]

	array := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice = array[:]
	fmt.Printf("len:%d, cap:%d\n", len(slice), cap(slice))
	fmt.Println(slice)
	slice = array[4:8]
	fmt.Printf("len:%d, cap:%d\n", len(slice), cap(slice))
	fmt.Println(slice)

	//也可以直接new一个数组给切片(slice = [5]int{2,3,4,5,6})这样是不允许的
	slice = []int{2, 3, 4, 5, 6}
	fmt.Printf("len:%d, cap:%d\n", len(slice), cap(slice))
	fmt.Println(slice)

	//切片也可以通过make函数创建 make([]类型,长度,容量) 容量可以不写,容量不写就默认等于长度
	slice = make([]int, 7, 7)
	fmt.Printf("len:%d, cap:%d\n", len(slice), cap(slice))
	fmt.Println(slice)

	slice = make([]int, 7, 10)
	fmt.Printf("len:%d, cap:%d\n", len(slice), cap(slice))
	fmt.Println(slice)

	//切片是不能直接比较的 (切片不能比较会直接提示语法错误的)
	// var slice1 []int = []int{2, 3, 4, 5, 6}
	// var slice2 []int = []int{2, 3, 4, 5, 6}
	// if slice1 != slice2 {
	// 	fmt.Println("切片相等")
	// }

	//切片追加元素
	slice1 := []int{1, 2, 3}
	slice2 := []int{14, 15, 16}
	slice1 = append(slice1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Printf("len:%d, cap:%d\n", len(slice1), cap(slice1))
	fmt.Println(slice1)
	//切片追加切片元素
	slice1 = append(slice1, slice2...)
	fmt.Printf("len:%d, cap:%d\n", len(slice1), cap(slice1))
	fmt.Println(slice1)

	//拷贝切片
	var slice3 []int
	var slice4 []int = []int{1, 2, 3, 4, 5, 6}
	slice3 = []int{2, 3, 4}
	copy(slice3, slice4)
	fmt.Println(slice3)
	fmt.Println(slice4)
}
