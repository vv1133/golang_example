package main

import (
	"fmt"
)

type testStruct1 struct {
	i int
	s string
}

func main() {
	var ts1 *testStruct1 = &testStruct1{}
	//var ts2 **testStruct1 = &ts1
	//var ts1 *testStruct1 = new(testStruct1)
	//ts1 := new(testStruct1)
	//ts2.i = 4
	//ts2.s = "aaa"
	(*ts1).i = 4
	(*ts1).s = "aaa"
	fmt.Println(ts1)
	//fmt.Println(*ts2)
}
