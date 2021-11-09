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

	H, W := nextInt(), nextInt()

	S := make([][]string, H)
	for i := 0; i < H; i++ {
		S[i] = strings.Split(nextString(), "")
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[i][j] == "." {
				S[i][j] = "0"
			}
			if S[i][j] == "#" {
				S[i][j] = "X"
			}
		}
	}
	for _, v := range S {
		fmt.Println(strings.Join(v, ""))
	}
}
