package main

import (
	"fmt"
)

func main() {
	var s string
	s = "hello world"
	bytes := make([]byte, 10)
	bytes = []byte(s)
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}
