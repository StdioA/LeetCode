package main

func genSequence(n int) []int {
	var seq = []int{0, 1}
	for i := 1; i < n; i++ {
		var (
			layer = make([]int, len(seq)*2)
			l     = len(layer)
		)

		// 0 正 1 反
		for j, n := range seq {
			layer[j], layer[l-1-j] = n, n^(1<<i)
		}
		seq = layer
	}
	return seq
}

func circularPermutationFirst(n int, start int) []int {
	// generate sequence
	seq := genSequence(n)
	for i, num := range seq {
		if num == start {
			return append(seq[i:], seq[:i]...)
		}
	}
	return seq
}

// 优化版本
func circularPermutationSecond(n int, start int) []int {
	var seq = []int{start, start ^ 1}
	for i := 1; i < n; i++ {
		// 0 正 1 反
		for j := len(seq) - 1; j >= 0; j-- {
			seq = append(seq, seq[j]^(1<<i))
		}
	}
	return seq
}

// 格雷码
func circularPermutation(n int, start int) []int {
	var seq = make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		seq[i] = start ^ i ^ (i >> 1)
	}
	return seq
}
