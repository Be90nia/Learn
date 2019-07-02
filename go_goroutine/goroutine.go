package main

import (
	"fmt"
	"runtime"
	"time"
)

//go 加上函数或者匿名函数就能开辟线程
//协程 Coroutine
//轻量级"线程"
//非抢占式多任务处理,由协程主动交出控制权
//编译器/解释器/虚拟机层面的多任务
//多个协程可能在一个或多个线程上运行
func main() {
	//for i := 0; i < 1000; i++ {
	//	go func(i int) {
	//		for {
	//			fmt.Printf("Hello from "+
	//			"goroutine %d\n", i)
	//		}
	//	}(i)
	//}
	//time.Sleep(time.Millisecond)

	//下面的代码是死循环
	//var a [10]int
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		for {
	//			//这样写没有交出控制权所以死循环在里面
	//			a[i]++
	//		}
	//	}(i)
	//}
	//time.Sleep(time.Millisecond)

	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//这样写没有交出控制权所以死循环在里面
				a[i]++
				//手动交出控制权让别的程序可以执行代码
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
