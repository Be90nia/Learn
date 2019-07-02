package main

import (
	"fmt"
	"time"
)
//chan
//创建chan的方法
//c := make(chan int)
//chan收发数据
//c<- 1 为发数据
//n := <- c 为收数据
//(注意chan的发数据后,必须要有收数据,否则程序会报错)
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
	for n := range c{
		fmt.Printf("Worker: %d received: %c\n",
				id, n)
	}
}
// chan<- int 属于发数据,不能收数据
// <-chan int 属于收数据,不能发数据
// 有点类似于get和set
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker: %d received: %c\n",
				id, <-c)
		}
	}()
	return c
}
//chan 也可以创建缓冲区进行数据存放
//这样就不需要创建数据接收
func bufferedChannel() {
	c := make(chan int, 3) //chan的创建同时创建3个缓冲区
	go worker(0 ,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd' //由于第四个没有缓冲区以及接收方的,所以会报错
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3) //chan的创建同时创建3个缓冲区
	go worker(0 ,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd' //由于第四个没有缓冲区以及接收方的,所以会报错
	close(c) //关闭发送channel
	time.Sleep(time.Millisecond)
}
func chanDemo() {
	// all goroutines are asleep - deadlock!
	//c := make(chan int)
	//c <- 1
	//c <- 2
	//n := <-c
	//fmt.Println(n)

	//c := make(chan int)
	//go worker(0,c)
	//c <- 1
	//c <- 2
	//time.Sleep(time.Millisecond)

	//var channels [10]chan int
	//for i := 0; i < 10; i++ {
	//	channels[i] = make(chan int)
	//	go worker(i, channels[i])
	//}
	//for i := 0; i < 10; i++ {
	//	channels[i] <- 'a' + i
	//}
	//for i := 0; i < 10; i++ {
	//	channels[i] <- 'A' + i
	//}
	//time.Sleep(time.Millisecond)

	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
//理论基础: Communication Sequential Process (CSP)
//不要通过共享内存来通信;通过通信来共享内存
func main(){
	//chanDemo()
	//bufferedChannel()
	channelClose()
}