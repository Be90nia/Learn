package main

import "fmt"

func main() {
	var array1 [4]int = [4]int{1, 3}
	fmt.Println(len(array1))
	fmt.Println(array1)

	var array2 = []int{1, 2, 2, 5}
	fmt.Println(len(array2))
	fmt.Println(array2)

	array3 := []int{4, 4568, 25, 4, 52, 76}
	fmt.Println(len(array3))
	fmt.Println(array3)
	fmt.Println(array3[2])
	//指针数组
	var array *[4]int = new([4]int)
	for i := 0; i < len(array); i++ {
		array[i] = i
	}
	fmt.Println(len(array))
	fmt.Println(array)
	//指针传递,不会产生多余的内存
	var array4 = array
	fmt.Println(array4[2])
	fmt.Printf("%p\n", array)
	fmt.Printf("%p\n", array4)
	//引用传递,应该是开辟新内存传递进去
	var array5 = &array
	fmt.Printf("%p\n", array)
	fmt.Printf("%p\n", array5)
}
