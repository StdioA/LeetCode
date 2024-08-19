package main

func reverseNum(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}

func rotateArray(nums []int, k int) {
	// 有点考验智力
	k = k % len(nums) // 保证 k 不要越界

	reverseNum(nums)
	reverseNum(nums[:k])
	reverseNum(nums[k:])
}
