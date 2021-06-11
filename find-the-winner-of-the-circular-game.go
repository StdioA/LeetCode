package main

// 模拟解
func findTheWinnerSlow(n int, k int) int {
	var nums = make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	var pointer = n - 1
	for round := 1; round < n; round++ {
		count := 0
		for count < k {
			pointer = (pointer + 1) % n
			if nums[pointer] > 0 {
				count++
			}
		}
		nums[pointer] = 0
	}
	for _, n := range nums {
		if n > 0 {
			return n
		}
	}
	return 0
}

// 内存拷贝缩短队列
func findTheWinner(n int, k int) int {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	var i = 0
	for len(nums) > 1 {
		i = (i + k - 1) % len(nums)
		// Remove element
		nums = append(nums[:i], nums[i+1:]...)
	}
	return nums[0]
}
