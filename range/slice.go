package main

import (
	"fmt"
)

func main() {
	str := [][]string{
		{"111", "aaa"},
		{"222", "bbb"},
		{"333", "ccc"},
	}

	for r, v := range str {
		fmt.Println(r, v)
		for c, vv := range v {
			fmt.Println(c, vv)
		}
	}
}
