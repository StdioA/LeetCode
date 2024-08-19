package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	var numbers = make([]string, len(nums))
	for i, num := range nums {
		numbers[i] = strconv.Itoa(num)
	}
	// 非常取巧的做法，直接把两个字符串拼起来，谁大谁就在前面
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i]+numbers[j] > numbers[j]+numbers[i]
	})

	result := strings.Join(numbers, "")
	// Remove leading zeroes
	for i, r := range result {
		if r != '0' {
			return result[i:]
		}
	}
	return "0"
}
