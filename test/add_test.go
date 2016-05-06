package hello

import "testing"

func TestAdd1(t *testing.T) {
	if Add(2, 4) == 6 {
		t.Log("test add1 passed")
	} else {
		t.Error("test add1 failed")
	}
}

func TestAdd2(t *testing.T) {
	if Add(1, 1) == 3 {
		t.Log("test add2 passed")
	} else {
		t.Error("test add2 failed")
	}
}
