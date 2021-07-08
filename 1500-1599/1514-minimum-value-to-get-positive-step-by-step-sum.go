package main

func minStartValue(nums []int) int {
	var (
		min, sum int
	)
	for _, num := range nums {
		sum += num
		if sum < min {
			min = sum
		}
	}
	if min < 0 {
		return -min + 1
	}
	return 1
}
