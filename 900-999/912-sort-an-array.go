package main

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	var (
		l     = len(nums)
		i, j  = 0, l - 1
		pivot = nums[l/2]
	)
	for i <= j {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	sortArray(nums[:j+1])
	sortArray(nums[i:])
	return nums
}
