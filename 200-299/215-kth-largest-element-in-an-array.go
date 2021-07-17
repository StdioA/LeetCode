package main

func findKthLargest(nums []int, k int) int {
	// 桶
	var (
		buckets = make([]int, 20001)
		count   int
	)
	for _, num := range nums {
		buckets[num+10000]++
	}

	for i := len(buckets) - 1; i >= 0; i-- {
		count += buckets[i]
		if count >= k {
			return i - 10000
		}
	}
	return 0
}
