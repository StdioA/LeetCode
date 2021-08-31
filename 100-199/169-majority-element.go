package main

func majorityElement(nums []int) int {
	var (
		mode    = nums[0]
		counter = 1
	)
	for _, num := range nums[1:] {
		if num == mode {
			counter++
		} else {
			counter--
		}
		if counter < 0 {
			mode = num
			counter = 1
		}
	}
	return mode
}
