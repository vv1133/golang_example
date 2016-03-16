package main

import (
	"fmt"
)

type Ele struct {
	interger int
}

func (ele Ele) Setx1(x int) {
	ele.interger = x
}

func (ele *Ele) Setx2(x int) {
	ele.interger = x
}

func main() {
	ele := new(Ele)
	ele.Setx1(2)
	fmt.Printf("%d\n", ele.interger)
	(*ele).Setx1(3)
	fmt.Printf("%d\n", ele.interger)
	ele.Setx2(4)
	fmt.Printf("%d\n", ele.interger)
	(*ele).Setx2(5)
	fmt.Printf("%d\n", ele.interger)
}
