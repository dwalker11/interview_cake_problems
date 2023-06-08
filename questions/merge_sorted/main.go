package main

import "fmt"

func mergeSortedArrays(a []int, b []int) []int {
	var sorted []int

	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if j == len(b) || a[i] < b[j] {
			sorted = append(sorted, a[i])
			i++
		} else if i == len(a) || a[i] > b[j] {
			sorted = append(sorted, b[j])
			j++
		}
	}

	return sorted
}

func main() {
	first := []int{3, 4, 6, 10, 11, 15}
	second := []int{1, 5, 8, 12, 14, 19}
	result := mergeSortedArrays(first, second)
	fmt.Println(result)
}
