package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var (
		result [][]int
		seen   = make(map[int]struct{})
	)
	if len(nums) == 0 {
		return result
	}

	sort.Ints(nums)
	for i, num := range nums {
		// Skip duplicate search for the first num
		if _, got := seen[num]; got {
			continue
		}

		sum := -num
		// search with double pointer
		begin, end := i+1, len(nums)-1
		for begin < end {
			b, e := nums[begin], nums[end]
			switch {
			case b+e == sum:
				result = append(result, []int{num, b, e})
				seen[num] = struct{}{}

				// Skip the duplicate
				for nums[begin] == b && begin < end {
					begin++
				}
				for nums[end] == e && begin < end {
					end--
				}
			case b+e < sum:
				begin++
			default:
				end--
			}
		}
		seen[num] = struct{}{}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{1, -1, -1, 0}))
}
