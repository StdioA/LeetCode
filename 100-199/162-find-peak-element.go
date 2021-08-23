package main

func findPeakElement(nums []int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)
	for left < right {
		mid := (left + right) / 2
		mid2 := mid + 1
		if nums[mid] < nums[mid2] {
			left = mid2
		} else {
			right = mid
		}
	}
	return left
}
