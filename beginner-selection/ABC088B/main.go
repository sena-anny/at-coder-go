package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func parseInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func parseToIntArray(str []string) []int {
	arr := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		n, _ := strconv.Atoi(str[i])
		arr[i] = n
	}
	return arr
}
func main() {
	line := nextLine()
	N := parseInt(line)
	line = nextLine()
	arr := parseToIntArray(strSplit(line))

	// int配列を逆順でソート
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	alice, bob := 0, 0
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			alice += arr[i]
		} else {
			bob += arr[i]
		}
	}
	result := alice - bob
	fmt.Println(result)

}
