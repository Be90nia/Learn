package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//案例1
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Printf("sum: %#v\n", sum)

	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(0),
	)

	printFile("abc.txt")
	forever()
}
func convertToBin(n int) string {
	result := ""
	if n == 0 {
		result = strconv.Itoa(n)
		return result
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//逐行读取文件
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//死循环
func forever() {
	i := 0
	for {
		fmt.Println("abc")
		i++
		if i == 5 {
			return //符合条件就跳出死循环
		}
	}
}
