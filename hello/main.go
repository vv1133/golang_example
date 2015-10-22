package main

import (
	"fmt"
)

var c = make(chan int, 10)

func f() {
	var a string
	a = "I'm f."
	fmt.Printf(a)
	c <- 0
}

func main() {
	fmt.Printf("hello world\n")
	go f()
	<-c
}
