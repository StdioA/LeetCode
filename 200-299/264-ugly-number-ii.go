package main

import "fmt"

// 答案是错的
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	var (
		nums = []int{
			2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25, 27, 30,
		}
		l   = len(nums)
		ans = 1
	)
	for i := 1; i <= (n-2)/l; i++ {
		ans *= 30
	}
	return ans * nums[(n-2)%l]
}

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Println(i, nthUglyNumber(i))
	}
}
