package main

func firstMissingPositive(nums []int) int {
	var (
		i      int
		length = len(nums)
	)
	for i < len(nums) {
		num := nums[i]
		if num > 0 && num <= length && nums[num-1] != num {
			// Place 1 to nums[0], 2 to nums[1], etc
			// if nums[0] is already filled with 1, then skip
			nums[num-1], nums[i] = nums[i], nums[num-1]
		} else {
			i++
		}
	}
	for i, n := range nums {
		if n != i+1 {
			return i + 1
		}
	}
	return length + 1
}
