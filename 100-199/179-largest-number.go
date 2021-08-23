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
