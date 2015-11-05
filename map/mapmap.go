package main

import "fmt"

func main() {
	map1 := make(map[string]map[string]int)

	if value, isPresent := map1["aaa"]; !isPresent {
		value = make(map[string]int)
		map1["aaa"] = value
	}
	map1["aaa"]["555"] = 5
	map1["aaa"]["444"] = 4
	fmt.Println(map1)
}
