package main

import (
	"fmt"
)

func tryRecover()  {
	defer func() {
		r := recover()
		if err,ok := r.(error); ok{
			fmt.Println("Error occurred:",err)
		}else {
			panic(fmt.Sprintf("I don't know what to do: %v",r))
		}
	}()
	//如果直接panic,defer中的recover会将错误进行处理
	//panic(errors.New("this is an error"))
	//b := 0
	//a := 5/b
	//fmt.Println(a)
	panic(123)
}
func main()  {
	tryRecover()
}
