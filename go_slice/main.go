package main

import "fmt"

//数组是传入函数中是以值传递的方式,如果想在函数中修改数组元素要么使用数组指针传递,或者是使用切片进行传递
func updateArray(array [8]int) {
	array[0] = 100
}

func updateSlice(array []int) {
	array[0] = 100
}

//切片
func main() {
	var slice []int
	fmt.Printf("len:%d, cap:%d\n", len(slice), cap(slice))

	//切片基础语法
	//切片初始化
	//切片名字 = 数组名[开始位置:结束位置]

	array := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	updateArray(array)
	fmt.Println("array", array)
	slice = array[:]

	updateSlice(slice)
	fmt.Println("slice", array)

	s1 := array[2:6] // 2,3,4,5
	fmt.Println("s1:", s1)
	s2 := s1[3:5] // 5,6 sile是可以向后扩展,但是不能超过数组长度,sile不能向前扩展
	fmt.Println("s2:", s2)

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

	delete_sile()
}

func delete_sile() {
	slice := []int{2, 3, 4, 5, 6}
	var slice1 = make([]int, 16)
	fmt.Println(slice1)
	copy(slice1, slice)
	fmt.Println("copy sile")
	fmt.Println(slice1)
	fmt.Println("deleting elements from slice")
	slice1 = append(slice1[:3], slice1[4:]...)
	fmt.Println(slice1)
}
