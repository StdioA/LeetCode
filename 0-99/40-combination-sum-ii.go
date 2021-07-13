package main

import "fmt"

func getSum(countMap map[int]int, target int) [][]int {
	if target == 0 {
		return [][]int{{}}
	}

	// Find num
	var (
		sum          int
		tempCountMap = make(map[int]int)
		targetNum    int
		targetCount  int
	)
	for num, count := range countMap {
		sum += num * count
		if count <= 0 {
			continue
		}
		if num > target {
			// 筛掉所有大于 target 的数
			tempCountMap[num] = count
		} else if targetNum == 0 {
			targetNum = num
			targetCount = count
			tempCountMap[num] = count
		}
	}

	// 现有总和小于 target，无解
	if sum < target {
		return nil
	}

	for num := range tempCountMap {
		countMap[num] = 0
	}
	defer func() {
		// 还原 count map
		for num, count := range tempCountMap {
			countMap[num] = count
		}
	}()

	var (
		result      = getSum(countMap, target)
		repeatedNum []int
	)

	for i := 1; i <= targetCount; i++ {
		repeatedNum = append(repeatedNum, targetNum)
		if targetNum*i > target {
			break
		}
		partial := getSum(countMap, target-targetNum*i)

		for _, partRes := range partial {
			result = append(result, append(partRes, repeatedNum...))
		}
	}

	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	var count = make(map[int]int)
	for _, num := range candidates {
		count[num]++
	}
	return getSum(count, target)
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
