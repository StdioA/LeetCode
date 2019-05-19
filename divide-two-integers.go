package main

import "fmt"

func divide(dividend int, divisor int) int {
    factor := false
    if dividend < 0 {
		factor = !factor
		dividend = -dividend
    }
    if divisor < 0 {
		factor = !factor
		divisor = -divisor
    }
    
    result := 0
    for dividend >= divisor {
		dividend -= divisor
        result++
    }
    if factor {
        result = -result
    }
    return result
}

func main() {
	fmt.Println(divide(10, 3))
}
