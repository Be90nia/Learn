package main

import (
	"fmt"
	"sync"
)
//chan
//创建chan的方法
//c := make(chan int)
//chan收发数据
//c<- 1 为发数据
//n := <- c 为收数据
//(注意chan的发数据后,必须要有收数据,否则程序会报错)
func doWork(id int, c chan int,done chan bool) {
	for n := range c {
		fmt.Printf("Worker: %d received: %c\n",
			id, n)
		done <- true
	}
}

func doWork2(id int, c chan int,wg *sync.WaitGroup ) {
	for n := range c {
		fmt.Printf("Worker: %d received: %c\n",
			id, n)
		wg.Done()
	}
}

func doWork3(id int, w worker3 ) {
	for n := range w.in {
		fmt.Printf("Worker: %d received: %c\n",
			id, n)
		w.done()
	}
}

type worker struct {
	in chan int
	done chan bool
}

type worker2 struct {
	in chan int
	wg *sync.WaitGroup
}

type worker3 struct {
	in chan int
	done func()
}

// chan<- int 属于发数据,不能收数据
// <-chan int 属于收数据,不能发数据
// 有点类似于get和set
func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}
func createWorker2(id int,wg *sync.WaitGroup) worker2 {
	w := worker2{
		in:   make(chan int),
		wg: wg,
	}
	go doWork2(id, w.in,wg)
	return w
}

func createWorker3(id int,wg *sync.WaitGroup) worker3 {
	w := worker3{
		in:   make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork3(id, w)
	return w
}
func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i,worker := range workers {
		worker.in <- 'a' + i
	}
	for _,worker := range workers {
		<-worker.done
	}
	for i,worker := range workers {
		worker.in <- 'A' + i
	}
	for _,worker := range workers {
		<-worker.done
	}

	var worker2s [10]worker2
	var wg  sync.WaitGroup
	for i := 0; i < 10; i++ {
		worker2s[i] = createWorker2(i, &wg)
	}
	wg.Add(20)
	for i,worker := range worker2s {
		worker.in <- 'a' + i
	}
	for i,worker := range worker2s {
		worker.in <- 'A' + i
	}
	wg.Wait()

	var worker3s [10]worker3
	var wg2  sync.WaitGroup
	for i := 0; i < 10; i++ {
		worker3s[i] = createWorker3(i, &wg2)
	}
	wg2.Add(20)
	for i,worker := range worker3s {
		worker.in <- 'a' + i
	}
	for i,worker := range worker3s {
		worker.in <- 'A' + i
	}
	wg2.Wait()
}
//理论基础: Communication Sequential Process (CSP)
//不要通过共享内存来通信;通过通信来共享内存
func main(){
	chanDemo()
}