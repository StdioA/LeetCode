package main

import "fmt"

func subsetsWithDup(nums []int) [][]int {
	var (
		freq   = make([]int, 21)
		result [][]int
	)
	for _, num := range nums {
		freq[num+10]++
	}
	// Calculate effective nums
	var effNums []int
	for i, f := range freq {
		if f > 0 {
			effNums = append(effNums, i-10)
		}
	}
	for i := 0; i < 1<<len(effNums); i++ {
		var (
			t      = i
			j      int
			layers = [][]int{{}}
		)
		for t > 0 {
			if t&1 == 0 {
				j++
				t >>= 1
				continue
			}
			// Copy & append nums with 1..n times
			var (
				num      = effNums[j]
				freq     = freq[num+10]
				temp     = [][]int{}
				multiNum = make([]int, 0, freq)
			)
			for k := 1; k <= freq; k++ {
				multiNum = append(multiNum, num)
				for _, res := range layers {
					// Becareful for space reusing on slices
					copied := make([]int, len(res), len(res)+k)
					copy(copied, res)
					temp = append(temp, append(copied, multiNum...))
				}
			}
			layers = temp
			j++
			t >>= 1
		}
		result = append(result, layers...)
	}
	return result
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 4, 3, 5, 4, 4, 7, 7, 8, 9, 0}))
}
