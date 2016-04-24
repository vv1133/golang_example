package main

import (
	"fmt"
	"github.com/vv1133/golang_example/doc/hello"
)

// Function main.
func main() {
	cat := hello.Cat{}
	cat.SetAge(2)
	dog := hello.Dog{}
	dog.SetAge(4)

	animals := []hello.Animal{&cat, &dog}
	for _, animal := range animals {
		fmt.Printf("%d\n", animal.Age())
	}
}
