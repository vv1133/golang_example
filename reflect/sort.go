package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

type T struct {
	B int
	C int
	A int
	x int
}

func (t T) String() string {
	return fmt.Sprintf("{%d %d}", t.A, t.B)
}

func main() {
	t := T{B: 2, A: 1, C: 3}
	fmt.Println(t)
	fmt.Println(&t)

	fmt.Println(printFields(t))
}

func printFields(st interface{}) string {
	t := reflect.TypeOf(st)

	names := make([]string, t.NumField())
	for i := range names {
		names[i] = t.Field(i).Name
	}
	sort.Strings(names)

	v := reflect.ValueOf(st)
	buf := &bytes.Buffer{}
	buf.WriteString("{")
	for i, name := range names {
		val := v.FieldByName(name)
		if !val.CanInterface() {
			continue
		}
		if i > 0 {
			buf.WriteString(" ")
		}
		fmt.Fprintf(buf, "%v", val.Interface())
	}
	buf.WriteString("}")

	return buf.String()
}
