package main

import (
	"fmt"
	"math"
)

// 完全背包
func numSquares(n int) int {
	var dp = make([][]int, 2)
	for i := 0; i <= 1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}
	dp[1] = dp[0][:]

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		value := i * i
		for j := value; j <= n; j++ {
			if dp[1][j] > dp[0][j-value]+1 {
				dp[1][j] = dp[0][j-value] + 1
			}
		}
		dp[0] = dp[1][:]
	}
	return dp[1][n]
}

// 勒让德三平方定理
func numSquares2(n int) int {
	q := int(math.Sqrt(float64(n)))
	if q*q == n {
		return 1
	}

	for n%4 == 0 {
		n /= 4
	}
	if n%8 == 7 {
		return 4
	}

	for i := 0; i*i <= n; i++ {
		j := int(math.Sqrt(float64(n - i*i)))
		if i*i+j*j == n {
			return 2
		}
	}
	return 3
}

func main() {
	fmt.Println(numSquares(13))
}
