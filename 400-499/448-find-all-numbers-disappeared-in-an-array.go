package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	abs := func(n int) int {
		if n < 0 {
			n = -n
		}
		return n
	}

	for _, n := range nums {
		index := abs(n) - 1
		nums[index] = -abs(nums[index])
	}
	var result []int
	for i, n := range nums {
		if n > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func main() {
	l := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(l))
}
