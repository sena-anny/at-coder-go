package main

import (
	"fmt"
	"time"
)

func Fib(number float64, ch chan float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	ch <- x
}

func main() {
	start := time.Now()

	size := 15
	ch := make(chan float64, size)
	for i := 0; i < size; i++ {
		go Fib(float64(i), ch)
		fmt.Printf("Fib(%v): %v\n", i, <-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
