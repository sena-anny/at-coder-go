package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func parseInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func main() {
	line := nextLine()
	A := parseInt(line)
	line = nextLine()
	B := parseInt(line)
	line = nextLine()
	C := parseInt(line)
	line = nextLine()
	X := parseInt(line)

	count := 0
	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				if i*500+j*100+k*50 == X {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
