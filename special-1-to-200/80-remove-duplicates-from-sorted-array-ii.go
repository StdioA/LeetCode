package main

func removeDuplicatesII(nums []int) int {
	// 裸题，与 1 差不多，不过边界条件需要想清楚
	if len(nums) == 0 {
		return 0
	}
	var (
		count       = 1
		assignIndex int // 最近一次赋过值的坐标，第一个元素已经自动算入了
	)
	for _, value := range nums[1:] {
		if value == nums[assignIndex] {
			// 相同的话 count++，前两个需要赋值，后面直接跳过
			count++
			if count <= 2 {
				assignIndex++
				nums[assignIndex] = value
			}
		} else {
			// 新元素与当前元素不同，则要赋新值，同时重置 count
			assignIndex++
			nums[assignIndex] = value
			count = 1
		}
	}
	return assignIndex + 1
}
