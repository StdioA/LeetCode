package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	count := make(map[int][]int)
	for i, num := range nums {
		count[num] = append(count[num], i)
	}

	Check := func(a, b, j int) bool {
		for _, k := range count[-a-b] {
			if k > j {
				return true
			}
		}
		return false
	}
	for i, a := range nums {
		for j, b := range nums[i+1:] {
			j += i + 1
			if Check(a, b, j) {
				result = append(result, []int{a, b, -a - b})
			}
		}
	}

	// remove duplicate result
	if len(result) == 0 {
		return result
	}

	SliceEqual := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}

	final := make([][]int, 0)
	for _, arr := range result {
		sort.Ints(arr)
	}
	for i := 0; i < len(result); i++ {
		f := true
		for j := 0; j < i; j++ {
			if SliceEqual(result[i], result[j]) {
				f = false
				break
			}
		}
		if f {
			final = append(final, result[i])
		}
	}
	return final
}

func main() {
	// l := []int{-1, 0, 1, 2, -1, -4}
	l := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(l))
}
