package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	t := make([]int, N+1)
	x := make([]int, N+1)
	y := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scan(&t[i], &x[i], &y[i])
	}

	flag := true
	for i := 1; i <= N; i++ {
		dt := int(math.Abs(float64(t[i] - t[i-1])))
		dx := int(math.Abs(float64(x[i] - x[i-1])))
		dy := int(math.Abs(float64(y[i] - y[i-1])))
		if dx+dy > dt || (dt-dx-dy)%2 != 0 {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
