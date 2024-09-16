package main

func findSmallestInteger(nums []int, value int) int {
	var modCount = make([]int, value)
	for _, num := range nums {
		mod := num % value
		if mod < 0 {
			mod += value
		}
		modCount[mod]++
	}
	// 找到最少出现的那个数，如果次数相同，则数字越小越好
	// 答案就是 value * c + n
	var (
		minimumCount = int(1e8)
		mininumNum   = 0
	)
	for num, count := range modCount {
		if count < minimumCount {
			minimumCount = count
			mininumNum = num
		} else if count == minimumCount && num < mininumNum {
			mininumNum = num
		}
	}
	return value*minimumCount + mininumNum
}
