package main

import "fmt"

func main() {
	var a string
	cnt := 0
	fmt.Scanf("%s", &a)
	if a[0] == '1' {
		cnt++
	}
	if a[1] == '1' {
		cnt++
	}
	if a[2] == '1' {
		cnt++
	}
	fmt.Print(cnt)
}
