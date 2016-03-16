package main

import (
	"fmt"
)

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	//names := []string{"stanley", "david", "oscar"}
	//PrintAll(names)//cannot use names (type []string) as type []interface {} in argument to PrintAll

	names := []string{"stanley", "david", "oscar"}
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)
}
