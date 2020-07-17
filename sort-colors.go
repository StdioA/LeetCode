package main

import "fmt"

func sortColors(nums []int) {
	start, end := 0, len(nums)-1
	i := 0
	for i <= end {
		c := nums[i]
		switch c {
		case 0:
			if i > start {
				nums[i], nums[start] = nums[start], nums[i]
			}
			start++
			if nums[i] >= 0 {
				i++
			}
		case 1:
			i++
		case 2:
			nums[i], nums[end] = nums[end], nums[i]
			end--
		}
	}
}

func main() {
	// h := []int{2,0,2,1,1,0}
	h := []int{2, 0, 1}
	sortColors(h)
	fmt.Println(h)
}
