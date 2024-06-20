package main

func containsNearbyDuplicate(nums []int, k int) bool {
	var lastIndex = make(map[int]int)
	// 既然是绝对值，那么只找到一组即可，因为 (i, j) 和 (j, i) 不一样
	// 第一个用例过于奇葩
	for i, num := range nums {
		if _, ok := lastIndex[num]; !ok {
			lastIndex[num] = i
			continue
		}
		li := lastIndex[num]
		if i-li <= k {
			return true
		}
		lastIndex[num] = i
	}
	return false
}

func main() {
	println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}
