package main

import "fmt"

var (
	freq    map[int]int
	curFreq map[int]int
	result  [][]int
)

func genPermutation(nums, current []int) {
	if len(current) == len(nums) {
		result = append(result, current[:])
		return
	}

	localFreq := make(map[int]bool)
	for _, num := range nums {
		// Check num is in current
		if curFreq[num] < freq[num] && !localFreq[num] {
			localFreq[num] = true
			curFreq[num]++
			genPermutation(nums, append(current, num))
			curFreq[num]--
		}
	}
}

func permuteUnique(arr []int) [][]int {
	result = nil
	curFreq = make(map[int]int)
	freq = make(map[int]int)
	for _, n := range arr {
		freq[n]++
	}
	genPermutation(arr, []int{})
	return result
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2, 1, 2}))
}
