package main

import "fmt"

// Time Complicity: O(N^2)
// Space Complicity: O(N)
func subarraySumSlow(nums []int, k int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	sum := make([]int, length)
	count := 0
	sum[0] = nums[0]
	if sum[0] == k {
		count++
	}

	for i := 1; i < length; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				sum[j] = sum[j] + nums[i]
			} else {
				sum[j] = sum[j-1] - nums[j-1]
			}
			if sum[j] == k {
				count++
			}
		}
	}
	return count
}

// Time Complicity: O(N)
// Space Complicity: O(N)
func subarraySum(nums []int, k int) int {
	sum, count := 0, 0
	sumMap := map[int]int{0: 1}
	for _, num := range nums {
		sum += num
		count += sumMap[sum-k]
		sumMap[sum]++
	}
	return count
}

func main() {
	nums := []int{3, 4, 7, 2, -3, 1, 4, 2}
	fmt.Println(subarraySum(nums, 7))
}
