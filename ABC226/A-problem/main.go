package main

import (
	"fmt"
	"math"
)

func main() {
	var X float64
	fmt.Scan(&X)

	fmt.Println(int(math.Round(X)))

}
