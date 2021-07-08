package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	var (
		bucket   = make([]int, 20001)
		max      = -20000
		maxIndex int
		result   = make([]int, 0, len(nums)-k+1)
	)
	const indexOffset = 10000
	// Build window
	for i := 0; i < k; i++ {
		num := nums[i]
		bucket[num+indexOffset]++
		if num > max {
			max = num
		}
	}
	maxIndex = max + indexOffset
	result = append(result, max)
	for i := k; i < len(nums); i++ {
		in, out := nums[i], nums[i-k]
		// Push in
		bucket[in+indexOffset]++
		// Pop out
		bucket[out+indexOffset]--

		if in > max {
			max = in
			maxIndex = max + indexOffset
		} else if out == max && bucket[maxIndex] == 0 {
			j := maxIndex
			for ; bucket[j] == 0; j-- {
			}
			max, maxIndex = j-indexOffset, j
		}
		result = append(result, max)
	}
	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
