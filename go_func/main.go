package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {

	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Println(q, r)
	//函数的第一个参数为函数名,因此go可以将函数当参数传递
	fmt.Println(apply(pow, 3, 4))
	//匿名函数
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	case "%":
		_, r := div(a, b)
		return r, nil
	default: //为了避免程序跳出运行,可以使用多返回值当函数出错就返回错误
		// panic("unsupported operation:" + op)
		return 0, fmt.Errorf("unsupported operation:" + op)
	}
}

//多返回函数
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d,%d)", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	sum := 0
	for i := range numbers { //类似foreach
		sum += numbers[i]
	}
	return sum
}
