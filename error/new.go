package main

//import "errors"
import "fmt"

func divideCheck(arg1, arg2 int) (error) {
	if arg2 == 0 {
		//return errors.New("can't divided by 0")
		return fmt.Errorf("%d can't divided by 0", arg1)
	}
	return nil
}

func main() {
	var err error

	err = divideCheck(4, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = divideCheck(8, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
}
