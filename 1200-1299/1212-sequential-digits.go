package main

import (
	"fmt"
	"math"
)

func genSeq(n, low, high int) []int {
	var (
		num  = 0
		step = 0
		seq  []int
	)
	for i := 1; i <= n; i++ {
		num = num*10 + i
		step = step*10 + 1
	}
	for i := 1; i <= 10-n; i++ {
		if num >= low && num <= high {
			seq = append(seq, num)
		}
		num += step
	}
	return seq
}

func sequentialDigits(low int, high int) []int {
	var (
		lowN   = int(math.Log10(float64(low))) + 1
		highN  = int(math.Log10(float64(high))) + 1
		result []int
	)
	for i := lowN; i <= highN; i++ {
		result = append(result, genSeq(i, low, high)...)
	}
	return result
}

func main() {
	fmt.Println(sequentialDigits(100, 300))
	fmt.Println(sequentialDigits(1000, 13000))
}
