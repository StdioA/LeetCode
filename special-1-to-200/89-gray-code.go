package main

func grayCode(n int) []int {
	// 格雷码
	var result = []int{0, 1}
	for i := 1; i < n; i++ {
		// 原样复制 x
		newResult := result
		// 倒序复制 x，并在高位填上 1
		for j := len(result) - 1; j >= 0; j-- {
			newResult = append(newResult, result[j]|(1<<i))
		}
		result = newResult
	}
	return result
}

func grayCode2(n int) []int {
	// 这个递推公式还是挺难想的
	var result = make([]int, 1<<n)
	for i := range result {
		result[i] = i ^ i>>1
	}
	return result
}
