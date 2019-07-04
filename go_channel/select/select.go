package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	//for {
	//	//由于关闭chan 时还会接收到chan的信息
	//	//因此需要进行chan关闭的跳出循环的操作
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker: %d received: %c\n",
	//		id, n)
	//}

	//此方法与上面的方法一样
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker: %d received: %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id,c)
	return c
}
//通过chan和select的方式可以创建一个非阻塞
func main() {
	//var c1,c2 chan int //c1 and c2 = nil
	//select {
	//case n:= <-c1:
	//	fmt.Println("Received form c1:",n)
	//case n:= <-c2:
	//	fmt.Println("Received form c2:",n)
	//default: //nil
	//	fmt.Println("No value received")
	//}

	//var c1, c2= generator(), generator()
	//for{
	//	select {
	//	case n := <-c1:
	//		fmt.Println("Received form c1:", n)
	//	case n := <-c2:
	//		fmt.Println("Received form c2:", n)
	//	}
	//}

	//var c1, c2= generator(), generator()
	//worker := createWorker(0)
	//n := 0
	//hasValue := false
	//for{
	//	var activeWorker chan<- int
	//	if hasValue{
	//		activeWorker = worker
	//	}
	//	select {
	//	case n = <-c1:
	//		hasValue = true
	//	case n = <-c2:
	//		hasValue = true
	//	case activeWorker <- n:
	//		hasValue = false
	//	}
	//}

	//csp模型
	var c1, c2= generator(), generator()
	worker := createWorker(0)

	var values []int
	tm := time.After(10*time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)//存储缓存数据
		case n := <-c2:
			values = append(values, n)//存储缓存数据
		case activeWorker <- activeValue:
			values = values[1:] //取出缓存的任务数据
		case <-time.After(800 * time.Millisecond)://超时处理
			fmt.Println("timeout")
		case <-tick://缓存数据长度检查
			fmt.Printf("queue len = %d\n",len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}


		//传统同步机制
		//WaitGroup
		//Mutex
		//Cond
	}
}
