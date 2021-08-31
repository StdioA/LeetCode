package main

import "fmt"

func firstMissingPositive(nums []int) int {
	var (
		i      int
		length = len(nums)
	)
	for i < len(nums) {
		num := nums[i]
		if num > 0 && num <= length && nums[num-1] != num {
			// Place 1 to nums[0], 2 to nums[1], etc
			// if nums[0] is already filled with 1, then skip
			nums[num-1], nums[i] = nums[i], nums[num-1]
		} else {
			i++
		}
	}
	for i, n := range nums {
		if n != i+1 {
			return i + 1
		}
	}
	return length + 1
}

// https://leetcode.com/problems/first-missing-positive/discuss/17080/Python-O(1)-space-O(n)-time-solution-with-explanation
// 问题的关键在于如何在 a[i-1] 记录 i 出现的频率，同时还要保证 a[i-1] 原本的数字能被算出来
// 思路类似的解答：https://leetcode.com/problems/first-missing-positive/discuss/17214/Java-simple-solution-with-documentation
func firstMissingPositive2(nums []int) int {
	// [-1, 4, 2, 0, 1, 2, 10]

	// 1. Add a additional slot for nums[0], which is used in step 4
	// [-1, 4, 2, 0, 1, 2, 10, 0]
	nums = append(nums, 0)

	var l = len(nums)
	// 2. Remove the numbers out of range
	// [0, 4, 2, 0, 1, 2, 0, 0]
	for i, n := range nums {
		if n < 0 || n > l {
			nums[i] = 0
		}
	}
	// 3. Set the frequency by l
	// for every (i, number), num / l is the frequency for num i, and num % l is the original number
	// [0 + 4l, 4 + 1l, 2 + 2l, 0, 1 + 1L, 2, 0, 0]
	for _, num := range nums {
		nums[num%l] += l
	}
	// frequency: [4, 1, 2, 0, 1, 0, 0, 0]
	// 4. nums[0] will be the place for number 0 or (l), which is len(original_nums) + 1, and we should't count for it
	for i := 1; i < l; i++ {
		if nums[i]/l == 0 {
			return i
		}
	}
	return l
}

func main() {
	var arr = []int{-1, 4, 2, 1, 9, 10}
	fmt.Println(firstMissingPositive(arr))
}
