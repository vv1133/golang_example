package main

import (
	"fmt"
)

type testA struct {
	a int
	b int
}

type testB struct {
	c int
}

type testC struct {
	testA
	testB
	d int
}

func main() {
	test := testC{testA{1, 2}, testB{4}, 54}
	fmt.Println(test)
}
