package main

import "fmt"

func findDuplicate(nums []int) int {
	set := make(map[int]int)
	for _, num := range nums {
		_, ext := set[num]
		if ext {
			return num
		} else {
			set[num] = 1
		}
	}
	return 0
}

func findDuplicate2(nums []int) int {
	var (
		slow = nums[0]
		fast = nums[nums[0]]
	)
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	fmt.Println(findDuplicate2([]int{1, 3, 4, 2, 2}))
}
