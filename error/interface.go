package main

import "fmt"

type division struct {
	arg int
	str string
}

func (e *division) Error() string {
	return fmt.Sprintf("%d %s", e.arg, e.str)
}

func divideCheck(arg1, arg2 int) (error) {
	if arg2 == 0 {
		return &division{arg1, "can't divided by 0"}
	}
	return nil
}

func main() {
	if err := divideCheck(4, 2); err != nil {
		fmt.Println(err)
	}
	if err := divideCheck(8, 0); err != nil {
		fmt.Println(err)
	}
}
