package main

import (
	"fmt"
)

func fun(i int) {
	defer fmt.Println("defer 1: ", i)
	if i == 0 {
		return
	}

	fmt.Println("fun")

	defer fmt.Println("defer 2: ", i)
}

func main() {
	var i int
	i = 3
	fmt.Println("start ", i)
	fun(i)

	i = 0
	fmt.Println("start ", i)
	fun(i)
}
