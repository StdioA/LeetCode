package main

import (
	"fmt"
	"sort"
)

func sortedSquares(A []int) []int {
	results := make([]int, len(A))
	for i, n := range A {
		results[i] = n * n
	}
	sort.Ints(results)
	return results
}

func main() {
	a := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(a))
}
