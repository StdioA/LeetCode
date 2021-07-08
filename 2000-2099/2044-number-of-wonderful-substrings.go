package main

import "fmt"

func wonderfulSubstrings(word string) int64 {
	var (
		count = make([]int64, 1024) // count[x] 表示 x 的前缀字符串有多少组，如 x=1 时，可能有 a, abb, bab, acc 等
		cur   int
		total int64
	)
	// Empty string
	count[0] = 1
	for i := 0; i < len(word); i++ {
		bit := 1 << (word[i] - 'a')
		cur ^= bit
		total += count[cur] // mask = 0, x ^ cur = mask
		for i := 0; i < 10; i++ {
			mask := 1 << i
			total += count[cur^mask] // mask = 1<<n, 只有一个奇数，x ^ cur = mask
		}
		count[cur]++
	}
	return total
}

func main() {
	fmt.Println(wonderfulSubstrings("aba"))
}
