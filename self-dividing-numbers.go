package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)
	for i := left; i <= right; i++ {
		num := i
		temp := i
		for temp > 0 {
			k := temp % 10
			if k == 0 || k > 0 && num%(temp%10) != 0 {
				break
			}
			temp /= 10
		}
		if temp == 0 {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
