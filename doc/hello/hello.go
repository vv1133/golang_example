// Copyright 2016 vv.
// This should not be shown in godoc.

// This is an example project for godoc.
package hello

// Animal is the interface implemented by any
// value that has an Age() method.
type Animal interface {
	// Age returns the value of animal's age.
	Age() int
	// SetAge sets the age of this animal.
	SetAge(int)
}

// Cat infomation.
type Cat struct {
	age int
}

// Dog infomation.
type Dog struct {
	age int
}

func (this Cat) Age() int {
	return this.age
}

func (this *Cat) SetAge(age int) {
	this.age = age
}

func (this Dog) Age() int {
	return this.age
}

// BUG(r): Dog SetAge has no effect.

func (this Dog) SetAge(age int) {
	this.age = age
}

