package main

import (
	"fmt"
)

func Set10(x int) {
	x = 10
}

func Set11(x *int) {
	//x = 11 //compile error
	*x = 11
}

func main() {
	a := 4
	Set10(a)
	fmt.Printf("a = %d\n", a)
	//Set11(a) //runtime error
	Set11(&a)
	fmt.Printf("a = %d\n", a)
}
