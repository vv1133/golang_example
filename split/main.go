package main

import (
	"fmt"
	"strings"
)

func main() {
	tag := "aaa,bbb,ccc"
	splits := strings.Split(tag, ",")
	fmt.Println(splits)

	for a, b := range splits {
		fmt.Println(a, b)
	}
}
