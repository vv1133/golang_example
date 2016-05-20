package main

import (
	"fmt"
	"runtime/pprof"
	"log"
	"os"
	"flag"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	ch := make(chan int)
	go func () {
		sum := 0
		for i:=0; i<100; i++ {
			sum += Fib(30)
		}
		ch <- sum
	}()
	i := <- ch
	fmt.Printf("%d\n", i)
}
