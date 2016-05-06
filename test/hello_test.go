package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	Hello()
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello()
	}
}

func ExampleHello_test1() {
	Hello()
	// Output: hello world
}

func ExampleHello_test2() {
	Hello()
	// Output: hello
}

