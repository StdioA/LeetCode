package main

import "fmt"

func searchRangeStart(nums []int, target int) int {
	// 查找最左边界
	var start, end = 0, len(nums) - 1
	for start < end {
		mid := (start + end) / 2
		switch {
		case nums[start] == target:
			return start
		case nums[mid] < target:
			start = mid + 1
		case nums[mid] == target:
			end = mid
		case nums[mid] > target:
			end = mid - 1
		}
	}
	return start
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	start := searchRangeStart(nums, target)
	end := searchRangeStart(nums, target+1)
	fmt.Println(start, end)
	if nums[start] != target {
		start = -1
	}
	// 用 target+1 去查找右边界的时候，end 和 end-1 有可能都是正确答案，需要单独判断下
	switch {
	case nums[end] == target:
	case end > 0 && nums[end-1] == target:
		end--
	default:
		end = -1
	}
	return []int{start, end}
}
