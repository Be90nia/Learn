package main

import (
	"fmt"
	"unicode/utf8"
)

func StringText() {

	s := "Yes我爱!" // UTF-8
	for _, b := range []byte(s) {
		fmt.Printf("%x", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d,%x)", i, ch)
	}
	fmt.Println()

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d,%c)", i, ch)
	}
}
