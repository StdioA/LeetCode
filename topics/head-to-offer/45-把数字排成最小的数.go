package main

import (
	"bytes"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		si, sj := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return si+sj < sj+si
	})
	var buf bytes.Buffer
	for _, n := range nums {
		buf.WriteString(strconv.Itoa(n))
	}
	return buf.String()
}
