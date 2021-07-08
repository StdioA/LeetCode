package main

import "fmt"

func permute2(nums []int) [][]int {
	length := len(nums)
	if length == 1 {
		return [][]int{nums}
	}

	result := make([][]int, 0)
	for i, n := range nums {
		x := make([]int, length-1, length-1)
		copy(x[:i], nums[:i])
		copy(x[i:], nums[i+1:])
		part := permute(x)
		for _, p := range part {
			result = append(result, append(p, n))
		}
	}
	return result
}

func genPermutation(nums, current []int, result [][]int) [][]int {
	if len(current) == len(nums) {
		result = append(result, current[:])
		return result
	}
	for _, num := range nums {
		// Check num is in current
		var occured bool
		for _, c := range current {
			if c == num {
				occured = true
				break
			}
		}
		if !occured {
			result = genPermutation(nums, append(current, num), result)
		}
	}
	return result
}

func permute(arr []int) [][]int {
	return genPermutation(arr, []int{}, [][]int{})
}

func main() {
	fmt.Println(permute2([]int{1, 2, 3, 4, 5}))
}
