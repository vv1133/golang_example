package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("abcd")
	buffer.WriteString("ffffff")
	fmt.Print(buffer.String(), "\n")

	var buffer2 bytes.Buffer
	buffer2.WriteString("1234")

	fmt.Print(buffer.String() + buffer2.String())
}
