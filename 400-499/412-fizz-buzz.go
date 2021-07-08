package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		ans[i-1] = s
	}
	return ans
}

func main() {
	fmt.Println(fizzBuzz(15))
}
