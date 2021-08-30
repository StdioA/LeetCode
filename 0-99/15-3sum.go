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

// 搜索剪枝，没懂它为什么比上面的那个要快
func threeSumFaster(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[j]+nums[k]+nums[i] > 0 {
				k--
			}
			if j == k {
				break
			}
			if nums[j]+nums[k]+nums[i] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}

func threeSumTLE(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	count := make(map[int][]int)
	for i, num := range nums {
		count[num] = append(count[num], i)
	}

	Check := func(a, b, j int) bool {
		for _, k := range count[-a-b] {
			if k > j {
				return true
			}
		}
		return false
	}
	for i, a := range nums {
		for j, b := range nums[i+1:] {
			j += i + 1
			if Check(a, b, j) {
				result = append(result, []int{a, b, -a - b})
			}
		}
	}

	// remove duplicate result
	if len(result) == 0 {
		return result
	}

	SliceEqual := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}

	final := make([][]int, 0)
	for _, arr := range result {
		sort.Ints(arr)
	}
	for i := 0; i < len(result); i++ {
		f := true
		for j := 0; j < i; j++ {
			if SliceEqual(result[i], result[j]) {
				f = false
				break
			}
		}
		if f {
			final = append(final, result[i])
		}
	}
	return final
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{1, -1, -1, 0}))
}
