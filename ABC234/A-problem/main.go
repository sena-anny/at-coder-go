package main

import "fmt"

func fx(x int) int {
	return x*x + 2*x + 3
}

func main() {
	var t int
	fmt.Scan(&t)

	fmt.Println(fx(fx(fx(t)+t) + fx(fx(t))))
}
