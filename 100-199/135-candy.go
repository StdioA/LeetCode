package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 复杂但正确
// 时间 96.43%，内存 7.14%
func candyComplicated(ratings []int) int {
	var (
		direction      int
		lowestIndices  []int
		highestIndices []int
		candies        = make([]int, len(ratings))
	)

	for i := range ratings {
		var delta int
		if i < len(ratings)-1 {
			delta = ratings[i+1] - ratings[i]
		}
		if delta > 0 && direction <= 0 {
			lowestIndices = append(lowestIndices, i)
		} else if delta < 0 && direction >= 0 {
			highestIndices = append(highestIndices, i)
		}
		if delta != 0 {
			direction = delta
		}
	}
	if direction < 0 {
		lowestIndices = append(lowestIndices, len(ratings)-1)
	} else if direction > 0 {
		highestIndices = append(highestIndices, len(ratings)-1)
	}

	for len(lowestIndices) > 0 && len(highestIndices) > 0 {
		lowestIndex, highestIndex := lowestIndices[0], highestIndices[0]
		count := 1
		if lowestIndex < highestIndex {
			// 从左往右发
			for j := lowestIndex; j <= highestIndex; j++ {
				if j > 0 && ratings[j] > ratings[j-1] {
					count++
				} else {
					count = 1
				}
				candies[j] = max(candies[j], count)
			}
			lowestIndices = lowestIndices[1:]
		} else {
			// 从右往左发
			for j := lowestIndex; j >= highestIndex; j-- {
				if j+1 < len(ratings) && ratings[j] > ratings[j+1] {
					count++
				} else {
					count = 1
				}
				candies[j] = max(candies[j], count)
			}
			highestIndices = highestIndices[1:]
		}
	}
	var total int
	for _, candy := range candies {
		if candy == 0 {
			candy = 1
		}
		total += candy
	}
	return total
}

// 找了个最快的解
func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	// 从左往右扫，累加上升值
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] += candies[i-1] + 1
		}
	}
	// 从右往左，累加上升值并跟之前的值取个 max
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}
	total := 0
	for _, candy := range candies {
		total += candy + 1
	}
	return total
}

func main() {
	fmt.Println(candy([]int{1, 1, 2, 2, 3, 2, 1, 1, 2, 3, 2, 4, 4}))
}
