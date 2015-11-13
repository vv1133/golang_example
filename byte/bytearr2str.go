package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	bytes := [16]byte{0x61, 2, 3, 4, 5, 6, 7}
	//bytes := []byte("1234567890abcdef")
	//bytes := []byte("1234567890")
	fmt.Printf("%T\n", bytes)

	slice := bytes[:]
	fmt.Println(slice)

	str := string(slice)
	fmt.Println(str)

	fmt.Println(hex.EncodeToString(bytes[:]))
}
