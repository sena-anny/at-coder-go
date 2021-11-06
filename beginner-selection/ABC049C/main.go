package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	words := [4]string{"dream", "dreamer", "erase", "eraser"}
	for i := 0; i < len(words); {
		if strings.HasSuffix(S, words[i]) {
			S = strings.TrimSuffix(S, words[i])
			i = 0
			continue
		}
		i++
	}

	if len(S) > 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
