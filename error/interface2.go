package main

//import "errors"
import "fmt"

type division struct {
	err error
}

func (this *division)DivideCheck(arg1, arg2 int) {
	if this.err != nil {
		return
	}
	if arg2 == 0 {
		//this.err = errors.New("can't divided by 0")
		this.err = fmt.Errorf("%d can't divided by 0", arg1)
		return
	}
}

func (this *division)Err() error {
	return this.err
}

func main() {
	d := new(division)
	d.DivideCheck(4, 2)
	d.DivideCheck(8, 0)
	if d.Err() != nil {
		fmt.Println(d.Err())
	}
}
