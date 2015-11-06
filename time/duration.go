package main

import (
	"fmt"
	"time"
)

func main() {
	seconds := 10
	fmt.Println(time.Duration(seconds) * time.Second)
	fmt.Println(time.Duration(10 * 100))
}
