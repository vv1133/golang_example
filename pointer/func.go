package main

import (
	"fmt"
)

func UpdateA(x int) {
	x = 5
}

func UpdateB(x *int) {
	//x = 5 //compile error
	*x = 5
}

func main() {
	a := 4
	UpdateA(a)
	fmt.Printf("a = %d\n", a)

	b := 4
	//UpdateB(b) //compile error
	UpdateB(&b)
	fmt.Printf("b = %d\n", b)
}
