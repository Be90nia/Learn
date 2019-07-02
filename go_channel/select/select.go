package main

import "fmt"

//通过chan和select的方式可以创建一个非阻塞
func main(){
	var c1,c2 chan int //c1 and c2 = nil
	select {
	case n:= <-c1:
		fmt.Println("Received form c1:",n)
	case n:= <-c2:
		fmt.Println("Received form c2:",n)
	default: //nil
		fmt.Println("No value received")
	}
}
