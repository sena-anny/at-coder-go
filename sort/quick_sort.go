package main

import "fmt"

func QuickSort(a []int) []int {
	//要素数が1以下になった場合はソート終了
	if len(a) < 2 {
		return a
	}
	// 先頭を基準値とする
	pivot := a[0]

	var left []int
	var right []int
	for _, value := range a[1:] {
		if value > pivot {
			right = append(right, value)
		} else {
			left = append(left, value)
		}
	}

	//再起的に処理を実行
	left = QuickSort(left)
	right = QuickSort(right)

	var sortedValue []int
	sortedValue = append(left, pivot)
	sortedValue = append(sortedValue, right...)
	return sortedValue
}

func main() {
	a := []int{6, 2, 1, 4, 5, 3, 9, 20, 3, 4, 16}
	fmt.Println(QuickSort(a))
}
