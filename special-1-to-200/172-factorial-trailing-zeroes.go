package main

func trailingZeroes(n int) int {
	// 只需统计 [1, n] 中，5 的个数
	// 想象一个金字塔
	// 第一层中，n 的个数是 n / 5
	// 第二层中，是 n / 25
	// 然后累加起来
	zeroCount := 0
	for n > 0 {
		n /= 5
		zeroCount += n
	}
	return zeroCount
}
