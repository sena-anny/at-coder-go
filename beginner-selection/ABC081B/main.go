package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func strSplit(str string) []string {
	cols := strings.Split(str, " ")
	return cols
}

func checkAllEven(a []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i]%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	line := nextLine()
	N := parseInt(line)
	line = nextLine()
	spl := strSplit(line)

	intArray := make([]int, N)

	for i := 0; i < N; i++ {
		intArray[i] = parseInt(spl[i])
	}

	count := 0
	for checkAllEven(intArray) {
		count++
		for i := 0; i < N; i++ {
			intArray[i] = intArray[i] / 2
		}
	}
	fmt.Println(count)

}
