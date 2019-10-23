package main

import "fmt"

func singleNumber(nums []int) int {
	var low, high int
	for _, num := range nums {
		low = (low ^ num) & ^high
		high = (high ^ num) & ^low
	}
	return low
}

func main() {
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
}
