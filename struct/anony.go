package main

import (
	"fmt"
)

type testA struct {
	a  int
	b  int
	cc string
}

type testB struct {
	a int
	c int
}

type testC struct {
	testA
	testB
}

func main() {
	testc := testC{}
	testc.testA.a = 4
	testc.testB.a = 2
	testc.c = 4
	fmt.Println(testc)
}
