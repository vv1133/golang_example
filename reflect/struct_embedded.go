package main

import (
	"fmt"
	"reflect"
)

type B struct {
	X string
	Y string
}

type D struct {
	B
	Z string
}

type C struct {
	X string
	Y string
	Z string
}

func DeepFields(iface interface{}) []reflect.Value {
	fields := make([]reflect.Value, 0)
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)

		//fmt.Println(v.Kind())
		//fmt.Println(v)
		switch v.Kind() {
		case reflect.Struct:
			fields = append(fields, DeepFields(v.Interface())...)
		default:
			fields = append(fields, v)
		}
	}

	//maps := make(map[string]interface{}, ift.NumField())
	//for i := range fields {
	//	fields[i]
	//}

	return fields
}

func main() {
	var i interface{}
	d := D{Z: "baz"}
	d.X = "foo"
	d.Y = "bar"
	i = d
	t := reflect.TypeOf(i)
	fmt.Println(t.NumField(), t.Field(0).Type)

	c := C{"foo", "bar", "baz"}

	fmt.Println(DeepFields(d))
	fmt.Println(DeepFields(c))

	fmt.Println(d)
	fmt.Println(d.X, d.Y, d.Z)
	dd := DeepFields(d)

	fmt.Printf("dd: %T\n", dd)
	fmt.Println(dd[0], dd[1], dd[2])
	fmt.Println(dd)
}
