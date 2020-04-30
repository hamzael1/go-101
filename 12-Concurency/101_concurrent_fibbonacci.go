package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Naive Fibonnaci
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

var wg sync.WaitGroup
func main() {
	start := time.Now()
	//fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1) // My computer has 4 cores, when I use more than 1 core I gained ~ two times the speed
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i:= 3; i < 50; i++ {
			fmt.Println("A", i, fib(i))
		}
	}()

	go func() {
		defer wg.Done()

		for i:= 3; i < 50; i++ {
			fmt.Println("B", i, fib(i))
		}
	}()

	wg.Wait()


	fmt.Println("Done", time.Since(start))
}
