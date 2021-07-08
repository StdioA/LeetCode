package main

import (
	"fmt"
)

func binaryGap(N int) int {
	max := 0
	dis := 0
	for N%2 == 0 {
		N >>= 1
	}

	for N > 0 {
		N >>= 1
		if N == 0 {
			break
		}

		dis = 1
		for N%2 == 0 && N > 0 {
			N >>= 1
			dis++
		}
		if dis > max {
			max = dis
		}
	}
	return max
}

func main() {
	fmt.Println(binaryGap(8))
	fmt.Println(binaryGap(22))
}
