package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(nextString())
	return i
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	N := nextInt()

	X := make([]int, N)
	Y := make([]int, N)

	for i := 0; i < N; i++ {
		X[i] = nextInt()
		Y[i] = nextInt()
	}
	var nowDistance float64
	var maxDistance float64
	maxDistance = 0
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			nowDistance = math.Sqrt(math.Pow(float64(X[i]-X[j]), 2) + math.Pow(float64(Y[i]-Y[j]), 2))
			if maxDistance < nowDistance {
				maxDistance = nowDistance
			}
		}
	}

	fmt.Println(maxDistance)
}
