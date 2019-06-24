package main

import "fmt"

func reverse(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		opp := len(nums) - 1 - i
		nums[i], nums[opp] = nums[opp], nums[i]
	}
}

func nextPermutation(nums []int) {
	length := len(nums)
	left, right := 0, length-1
	for right > 0 && nums[right] <= nums[right-1] {
		right--
	}
	if right > 0 {
		left = right - 1
		right = length - 1
		for right > left && nums[right] <= nums[left] {
			right--
		}
		if right > left {
			nums[left], nums[right] = nums[right], nums[left]
		}
		left++
	}
	// reverse
	reverse(nums[left:])
}

func main() {
	var l []int
	l = []int{1, 2}
	nextPermutation(l)
	fmt.Println(l)
	l = []int{1, 3, 2}
	nextPermutation(l)
	fmt.Println(l)
	l = []int{1, 2, 3}
	nextPermutation(l)
	fmt.Println(l)
}

// (l) / 2
// 3 -> i < 1
// 0 - 2
// 4 -> i < 2
// 0 - 3
// 5 -> i < 2
// 0 - 4
