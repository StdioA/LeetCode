package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	length := len(nums)
	sort.Ints(nums)
	max1 := nums[0] * nums[1] * nums[length-1]
	max2 := nums[length-1] * nums[length-2] * nums[length-3]
	if max1 < max2 {
		max1 = max2
	}
	return max1
}

func main() {
	list := []int{-4, -3, -2, -1, 60}
	fmt.Println(maximumProduct(list))
}
