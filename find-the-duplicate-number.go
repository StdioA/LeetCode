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

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}
