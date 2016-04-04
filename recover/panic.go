package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	fmt.Println("Calling g(0).")
	g(0)
	fmt.Println("Returned normally from g(0).")
}

func g(i int) {
	defer fmt.Println("Defer in g", i)

	fmt.Println("Panic!")
	panic("panic in g")
	fmt.Println("Printing in g", i)
}
