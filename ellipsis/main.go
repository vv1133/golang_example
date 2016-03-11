package main

import (
	"fmt"
)

func Greeting(prefix string, who ...string) {
	fmt.Println(prefix, who)

	for _, v := range who {
		fmt.Printf("%s\n", v)
	}
}

func main() {
	Greeting("nobody")
	Greeting("hi:", "OK", "WT")
	Greeting("hello:", "Joe", "Anna", "Eileen")

	s := []string{"James", "Jasmine"}
	Greeting("goodbye:", s...)
}
