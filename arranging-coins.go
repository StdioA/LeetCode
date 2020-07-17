package main

import "fmt"

func arrangeCoins(n int) int {
	var x int
	for i := 1; i <= n+1; i++ {
		x += i
		if x > n {
			return i - 1
		}
	}
	return 0
}

func main() {
	fmt.Println(arrangeCoins(1))
	fmt.Println(arrangeCoins(8))
	fmt.Println(arrangeCoins(10))
}
