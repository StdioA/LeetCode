package main

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	checkValid := func(a, b, c int, count map[int]int) bool {
		for _, num := range([]int{c, a, b}) {
			count[num]--
			if count[num] < 0 {
				return false
			}
		}
		return true
	}
	for a, _ := range count {
		for b, _ := range count {
			c := -a-b
			if !(a <= b && b <= c) {
				continue
			}
			// copy map
			tmpCount := map[int]int{
				a: count[a],
				b: count[b],
				c: count[c],
			}
			if checkValid(a, b, c, tmpCount) {
				result = append(result, []int{a, b, c})
			}
		}
	}
	
	return result
}

func main() {
	// l := []int{-1, 0, 1, 2, -1, -4}
	l := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(l))
}
