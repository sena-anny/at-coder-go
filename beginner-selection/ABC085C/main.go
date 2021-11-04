package main

import "fmt"

func main() {
	var N, Y int
	fmt.Scan(&N, &Y)

	x, y, z := -1, -1, -1
	for i := 0; i <= N; i++ {
		for j := 0; j <= N-i; j++ {
			if Y == 10000*i+5000*j+1000*(N-i-j) {
				x = i
				y = j
				z = N - i - j
				break
			}
		}
	}
	fmt.Printf("%d %d %d", x, y, z)
}
