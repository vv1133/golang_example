package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]int)
	map1["key1"] = 1
	map1["key2"] = 2
	map1["key3"] = 3

	for k, v := range map1 {
		fmt.Println(k, v)
	}
}
