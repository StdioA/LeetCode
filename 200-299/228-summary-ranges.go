package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	if len(nums) == 0 {
		return result
	}

	segments := make([][]int, 0)
	start := nums[0]
	end := nums[0]
	for _, num := range nums[1:] {
		if num == end+1 {
			end = num
		} else {
			segments = append(segments, []int{start, end})
			start = num
			end = num
		}
	}
	segments = append(segments, []int{start, end})

	for _, seg := range segments {
		start, end = seg[0], seg[1]
		if start == end {
			result = append(result, strconv.Itoa(start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, end))
		}
	}
	return result
}

func main() {
	l := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(l))
}
