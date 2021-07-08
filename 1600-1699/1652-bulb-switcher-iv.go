package main

// 000101
// 去除前缀 0 后，判断后面有多少位翻转
func minFlips(target string) int {
	var (
		count int
		valid bool
	)

	for i := range target {
		if target[i] == '1' {
			valid = true
		}
		if i == 0 {
			if target[i] == '1' {
				count++
			}
		} else if valid && target[i] != target[i-1] {
			count++
		}
	}
	return count
}
