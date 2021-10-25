package main

import "fmt"

func Merge(left, right []int) []int {
	var result []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return append(result, append(left, right...)...)
}

func MergeSort(a []int) []int {
	if len(a) > 1 {
		n := len(a) / 2
		a = Merge(MergeSort(a[:n]), MergeSort(a[n:]))
		fmt.Println(a)
	}
	return a
}

func main() {
	arr := []int{3, 4, 5, 1, 6, 2, 7, 8}
	fmt.Println(MergeSort(arr))
}
