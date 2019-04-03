package main

import "fmt"

func search(nums []int, target int) int {
	var left, right, middle int
	left, right = 0, len(nums)-1
	for left < right {
		middle = (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func main() {
	l := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(l, 9))
	fmt.Println(search(l, 6))
	fmt.Println(search(l, -1))
	fmt.Println(search(l, -100))
	fmt.Println(search(l, 12))
	fmt.Println(search(l, 1200))
}
