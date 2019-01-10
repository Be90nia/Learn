package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//if条件判断语句
	//if 条件语句 {
	//	代码块
	//}else if 条件语句{
	// 	代码块
	//}else{
	//	代码块
	//}
	//条件语句均可以使用 &&、||、!逻辑运算符

	//switch 第一种用法
	i := 1
	switch i {
	case 1:
		fmt.Printf("i = %#v\n", i)
	case 2:
		fmt.Printf("i = %#v\n", i)
	}
	//switch 第二种用法
	i = 3
	switch {
	case i < 2:
		fmt.Printf("i < 2\n")
	case i < 4:
		fmt.Printf("i < 4\n")
	case i < 6:
		fmt.Printf("i < 6\n")
	}

	//go语言 只有一个for循环,没有while和do-while,for还有forech功能
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//如果for后面什么都不写，这样是默认为while(ture)的死循环的
	i = 0
	for {
		if i > 5 {
			break //跳出循环
		}
		i++
		if i == 3 {
			continue //跳出当前循环
		}
		fmt.Println(i)
	}

	array := [6]int{1, 2, 3, 4, 5, 6}
	for i, val := range array {
		fmt.Printf("%d: %d\n", i, val)
	}

	newmap := map[string]int{"one": 1, "two": 2}
	for i, val := range newmap {
		fmt.Printf("%s: %d\n", i, val)
	}
	var filname string = "abc.txt"
	if contents, err := ioutil.ReadFile(filname); err == nil {
		fmt.Println(string(contents))
	} else {
		fmt.Println("cannot print file contents:", err)
	}

}
