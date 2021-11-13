package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	A := map[string]bool{}
	for i := 0; i < N; i++ {
		L := nextInt()
		ss := make([]string, L)
		for j := 0; j < L; j++ {
			ss[j] = fmt.Sprintf("%011d", nextInt())
		}
		A[strings.Join(ss, "")] = true
	}
	for s, b := range A {
		fmt.Println(s, b)
	}
	fmt.Println(len(A))
}
