package main

func canJump(nums []int) bool {
	// 贪心题，尽量跳得更远
	var (
		index     int
		lastIndex = len(nums) - 1
	)

	for {
		if index+nums[index] >= lastIndex {
			return true
		}
		if nums[index] == 0 {
			break
		}
		var (
			nextJumpIndex = 0
			maxIndex      = index
		)
		for i := 1; i <= nums[index]; i++ {
			if nums[index+i]+index+i > maxIndex {
				maxIndex = nums[index+i] + index + i
				nextJumpIndex = index + i
			}
		}
		index = nextJumpIndex
	}
	return false
}
