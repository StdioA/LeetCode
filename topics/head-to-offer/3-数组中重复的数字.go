package main

func findRepeatNumber(nums []int) int {
	var i = 0
	for {
		num := nums[i]
		if num == i {
			i++
			continue
		}
		another := nums[num]
		if num == another {
			return num
		}
		nums[i], nums[num] = nums[num], nums[i]
	}
}
