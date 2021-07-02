package main

import "fmt"

// 时间复杂度 O(N)
// 空间复杂度 O(N)
func maxProduct(nums []int) int {
	var (
		length   = len(nums)
		positive = make([]int, length+1)
		negative = make([]int, length+1)
	)
	for i, num := range nums {
		if num == 0 {
			continue
		}
		pp, pn := positive[i], negative[i]
		if pp == 0 {
			pp = 1
		}
		if pn == 0 {
			pn = 1
		}

		if num > 0 {
			positive[i+1] = pp * num
			negative[i+1] = pn * num
			if num > positive[i+1] {
				positive[i+1] = num
			}
		} else {
			positive[i+1] = pn * num
			if num < negative[i+1] {
				negative[i+1] = num
			}
			negative[i+1] = pp * num
		}
	}
	// 从正数中找最大值，需要注意的是 positive 数组中的数字可能小于 0，如 nums = [-1] 时
	max := positive[1]
	for _, num := range positive[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

// 时间复杂度 O(N)
// 空间复杂度 O(1)
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProductLess(nums []int) int {
	var (
		curMin = nums[0]
		curMax = nums[0]
		result = nums[0]
	)
	for _, num := range nums[1:] {
		tmpMin := min(num, min(curMin*num, curMax*num))
		tmpMax := max(num, max(curMin*num, curMax*num))
		curMin = tmpMin
		curMax = tmpMax
		result = max(curMax, result)
	}
	return result
}

func main() {
	nums := []int{3, -1, 4}
	fmt.Println(maxProductLess(nums))
}
