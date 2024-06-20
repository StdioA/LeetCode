package main

// Note: 还有些原地修改 nums 的优化方法，但看起来太恶心了
func getMaximumXor(nums []int, maximumBit int) []int {
	var (
		sum     int
		length  = len(nums)
		answers = make([]int, length)
		mask    = 1<<maximumBit - 1
	)
	for _, num := range nums {
		sum ^= num
	}
	for i := len(nums) - 1; i >= 0; i-- {
		answers[length-i-1] = mask ^ sum
		sum ^= nums[i]
	}
	return answers
}
