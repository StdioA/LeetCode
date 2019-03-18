package main

import (
	"fmt"
	"math"
)

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	bits := int(math.Log2(float64(N))) + 1
	return N ^ (int(math.Pow(2, float64(bits))) - 1)
}

func main() {
	fmt.Println(bitwiseComplement(5))
}
