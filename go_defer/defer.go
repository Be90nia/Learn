package main

import (
	"bufio"
	"code/learn/go_functional/adder/fib"
	"fmt"
	"os"
)

//defer的调用
//确保调用在函数结束时发生

func tryDefer() {
	//此段代码输出的结果为 3 2 1
	//defer 的先后顺序为类似于遍历栈的方式,最后的defer运行,第一的defer最后运行(全部在函数return或者panic之前结束触发)
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)
	// return
	// fmt.Println(4)

	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)
	// panic("error occurred")
	// fmt.Println(4)
	for i := 0; i <100; i++{
		defer fmt.Println(i)
		if i == 30{
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	//创建一个文件
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()//函数运行完毕后释放内存

	writer := bufio.NewWriter(file) //创建文件写入
	defer writer.Flush()//函数运行完毕后将数据写入到文件中

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}

//defer 常见的defer调用
//Open/Close
//Lock/Unlock
//PrintHeader/PrintFooter