package main

import "fmt"

func findDuplicates(nums []int) []int {
	result := make([]int, 0)
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		index := num - 1
		nums[index] = -nums[index]
		if nums[index] > 0 {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	l := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(l))
}
