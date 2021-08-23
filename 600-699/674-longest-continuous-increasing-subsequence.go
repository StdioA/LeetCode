package main

func findLengthOfLCIS(nums []int) int {
	var (
		maxLength int
		length    = 1
	)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			if length > maxLength {
				maxLength = length
			}
			length = 1
		} else {
			length++
		}
	}
	if length > maxLength {
		maxLength = length
	}
	return maxLength
}
