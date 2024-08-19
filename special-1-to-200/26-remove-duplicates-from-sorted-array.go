package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var arrayTail int
	for _, num := range nums {
		if num != nums[arrayTail] {
			arrayTail++
			nums[arrayTail] = num
		}
	}
	return arrayTail + 1
}
