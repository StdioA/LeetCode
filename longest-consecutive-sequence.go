package main

import "fmt"

func longestConsecutiveSlow(nums []int) int {
	// 桶
	bucket := make(map[int]bool)
	for _, num := range nums {
		bucket[num] = true
	}
	max := 0
	for num := range bucket {
		count := 1
		for i := num - 1; bucket[i]; i-- {
			count++
		}
		if count > max {
			max = count
		}
	}
	return max
}

func longestConsecutive(nums []int) int {
	// 桶
	bucket := make(map[int]bool)
	for _, num := range nums {
		bucket[num] = true
	}
	// 减少重复查询
	max := 0
	for num := range bucket {
		count := 1
		for i := num - 1; bucket[i]; i-- {
			count++
			bucket[i] = false
		}
		for i := num + 1; bucket[i]; i++ {
			count++
			bucket[i] = false
		}
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
