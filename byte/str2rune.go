package main

import (
	"fmt"
)

func main() {
	s := "abcde好的"
	ru := []rune(s)
	bb := []byte(s)
	fmt.Println(len(ru), ru)
	fmt.Println(len(bb), bb)
}
