package main

func totalHammingDistance(nums []int) int {
	var (
		length = len(nums)
		total  int
	)
	for i := 0; i < 32; i++ {
		var ones int
		for _, num := range nums {
			ones += (num >> i) & 1
		}
		total += ones * (length - ones)
	}
	return total
}
