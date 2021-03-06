package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//斐波那契数列
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func fibonacci2() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// func main() {
// 	// f := Fibonacci()

// 	// fmt.Println(f())
// 	// fmt.Println(f())
// 	// fmt.Println(f())
// 	// fmt.Println(f())
// 	// fmt.Println(f())
// 	// fmt.Println(f())
// 	// fmt.Println(f())
// 	// fmt.Println(f())

// 	f := fibonacci2()
// 	printFileContents(f)
// }
