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

func strSplit(str string) []string {
	cols := strings.Split(str, " ")
	return cols
}

func parseToIntArray(str []string) []int {
	arr := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		n, _ := strconv.Atoi(str[i])
		arr[i] = n
	}
	return arr
}
func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}

func main() {
	line := nextLine()
	str := strSplit(line)
	intArray := parseToIntArray(str)

	N := intArray[0]
	A := intArray[1]
	B := intArray[2]

	total := 0
	for i := 1; i <= N; i++ {
		sum := sumOfDigits(i)

		if A <= sum && sum <= B {
			total = total + i
		}
	}
	fmt.Println(total)

}
