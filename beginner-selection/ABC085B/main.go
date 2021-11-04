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
	N := parseInt(line)
	d := make([]int, N)
	for i := range d {
		line = nextLine()
		d[i] = parseInt(line)
	}

	dm := make(map[int]bool)
	for _, value := range d {
		dm[value] = true
	}
	fmt.Println(len(dm))
}
