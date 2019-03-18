package main

import "fmt"
import "math"

func checkPerfectNumber(num int) bool {
	if num == 0 || num == 1 {
		return false
	}
	sum := 0
	for i := 1; i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			sum += i
			if i != 1 {
				sum += num / i
			}
		}
	}
	return sum == num
}

func main() {
	fmt.Println(checkPerfectNumber(1))
}
