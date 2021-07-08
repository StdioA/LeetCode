package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// find pivot
	l, r := 0, len(nums)-1
	for l != r {
		mid := (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	pivot := l

	// binary search
	if pivot == 0 {
		l, r = 0, len(nums)-1
	} else if target < nums[0] {
		l, r = pivot, len(nums)-1
	} else {
		l, r = 0, pivot-1
	}
	for (r - l) > 1 {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid
		} else {
			l = mid
		}
	}
	if nums[l] == target {
		return l
	} else if nums[r] == target {
		return r
	} else {
		return -1
	}
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 0, 1, 2}, 1))
}
