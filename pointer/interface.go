package main

import (
	"fmt"
)

type Animal interface {
	Age() int
}

type Cat struct {
	age int
}

func (cat Cat) Age() int {
	return cat.age
}

type Dog struct {
	age int
}

func (dog *Dog) Age() int {
	return dog.age
}

func main() {
	animals := []Animal{Cat{5}, &Dog{6}}
	for _, animal := range animals {
		fmt.Printf("%d\n", animal.Age())
	}
}
