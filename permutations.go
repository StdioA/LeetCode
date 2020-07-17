package main

import "fmt"

func permute(nums []int) [][]int {
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

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
