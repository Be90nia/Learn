package main

import "fmt"

func main() {
	//映射 可以类比成C#中的字典
	var mymap map[string]int //未被初始化的map是为nil的
	if mymap == nil {
		fmt.Println("map 为空的")
	}
	//make 也可用来初始化map
	mymap1 := make(map[string]int)
	fmt.Println(mymap1)

	mymap2 := make(map[string]int, 10)
	fmt.Println(mymap2)
	//map 赋值
	mymap3 := make(map[string]int)
	mymap3["one"] = 1
	fmt.Println(mymap3)
	mymap3 = map[string]int{"two": 2, "three": 3}
	mymap3["one"] = 1
	fmt.Println(mymap3)

	//访问映射
	fmt.Println(mymap3["one"])
	myvar, ok := mymap3["one"]
	fmt.Println(myvar, ok)
	myvar, ok = mymap3["four"] //参数名,参数名 = map[key] 第一个参数为映射值,第二个为key值是否存在(第一个参数可以使用占位符代替)
	fmt.Println(myvar, ok)
	_, ok = mymap3["four"] // _为占位符
	fmt.Println(ok)
	//映射删除
	delete(mymap3, "one")
	fmt.Println(mymap3)

	//映射的丰富值类型
	submap := make(map[string]int32)
	submap["one"] = 1
	submap["two"] = 2
	submap["three"] = 3
	mymap4 := make(map[string]map[string]int32)
	mymap4["map1"] = submap
	mymap4["map2"] = map[string]int32{"one": 1, "two": 2}
	fmt.Println(mymap4)
	fmt.Println(LengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(LengthOfNonRepeatingSubStr("是硕士双宿双飞"))

	StringText()
}
