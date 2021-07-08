package main

import "fmt"

func sortArrayByParity(A []int) []int {
	var result []int
	// for _, n := range A
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 1 {
			result = append(result, A[i])
		} else {
			result = append([]int{A[i]}, result...)
		}
	}
	return result
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}
