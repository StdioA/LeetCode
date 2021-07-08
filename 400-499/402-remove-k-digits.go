package main

import (
	"fmt"
)

func removeKdigits(nums string, k int) string {
	var (
		stack []rune
		count int
		i     int
		r     rune
		res   string
	)
	for i, r = range nums {
		stack = append(stack, r)
		for count < k && len(stack) > 1 {
			top := stack[len(stack)-1]
			if top >= stack[len(stack)-2] {
				break
			}
			stack[len(stack)-2] = top
			stack = stack[:len(stack)-1]
			count++
		}
		if count == k {
			break
		}
	}

	// strip zeros from stack
	for _, r := range stack {
		if r != '0' {
			break
		}
		stack = stack[1:]
	}

	if count < k {
		res = string(stack[:len(nums)-k])
	} else {
		res = string(stack)
		if i+1 < len(nums) {
			res += nums[i+1:]
		}
	}
	if len(res) == 0 {
		return "0"
	}
	return res
}

func main() {
	fmt.Println(removeKdigits("100", 1))
	input := "1234567890"
	for i := 1; i <= 9; i++ {
		fmt.Println(removeKdigits(input, i))
	}
}
