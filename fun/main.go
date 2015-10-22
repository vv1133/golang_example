package main

import (
	"fmt"
)

func main() {
	a := func() int {
		return 4
	}

	fmt.Println(a)
	fmt.Println(a())
}
