package main

import "fmt"

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var index int
	left, right := 0, 0
	// calculate right
	for _, num := range nums[1:] {
		right += num
	}
	for index = 0; index < len(nums)-1; index++ {
		if left == right {
			return index
		}
		left += nums[index]
		right -= nums[index+1]
	}
	if left == right {
		return index
	}
	return -1
}

func main() {
	l := []int{-1, -1, 0, 1, 1, 0}
	fmt.Println(pivotIndex(l))
}
