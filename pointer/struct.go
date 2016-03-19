package main

import (
	"fmt"
)

type Cat struct {
	age int
}

func (cat Cat) Age() int {
	return cat.age
}

func (cat Cat) SetAge(x int) {
	cat.age = x
}

type Dog struct {
	age int
}

func (dog Dog) Age() int {
	return dog.age
}

func (dog *Dog) SetAge(x int) {
	dog.age = x
}

func main() {
	cat := new(Cat)
	//cat := &Cat{1}
	//cat := Cat{age:1}

	cat.SetAge(3)
	fmt.Printf("%d\n", cat.Age())


	//dog := new(Dog)
	dog := &Dog{1}
	//dog := Dog{age:1}//compile error

	dog.SetAge(5)
	fmt.Printf("%d\n", dog.Age())
	fmt.Printf("%d\n", (*dog).Age())

	(*dog).SetAge(6)
	fmt.Printf("%d\n", dog.Age())
}
