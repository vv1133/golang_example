package main

import (
	"fmt"
)

type Shaper interface {
	Setx(x int)
	Getx() int
}

type EleA struct {
	x int
}

func (ele *EleA) Setx(x int) {
	ele.x = x
}

func (ele EleA) Getx() int {
	return ele.x
}

type EleB struct {
	x int
}

func (ele EleB) Setx(x int) {
	ele.x = x
}

func (ele EleB) Getx() int {
	return ele.x
}

func main() {
	var shape, shape2 Shaper

	//shape = EleA{5}//error
	shape = &EleA{5}
	fmt.Printf("%d\n", shape.Getx())
	shape.Setx(6)
	fmt.Printf("%d\n", shape.Getx())

	shape2 = EleB{5}
	fmt.Printf("%d\n", shape2.Getx())
	shape2.Setx(6)
	fmt.Printf("%d\n", shape2.Getx())

	//shape = &EleA{5}//ok
	//fmt.Printf("%d\n", shape.Getx())
	//shape.Setx(6)
	//fmt.Printf("%d\n", shape.Getx())

	//shape2 = &EleB{5}
	//fmt.Printf("%d\n", shape2.Getx())
	//shape2.Setx(6)
	//fmt.Printf("%d\n", shape2.Getx())
}
