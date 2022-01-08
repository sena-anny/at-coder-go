package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var K int64
	fmt.Scan(&K)

	bits := strconv.FormatInt(K, 2)
	fmt.Println(strings.ReplaceAll(bits, "1", "2"))
}
