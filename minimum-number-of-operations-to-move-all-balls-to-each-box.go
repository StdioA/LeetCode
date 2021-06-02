package main

import "fmt"

func dis(a, b int) int {
	var res = a - b
	if res < 0 {
		res = -res
	}
	return res
}

// O(N^2)
func minOperationsSlow(boxes string) []int {
	var (
		total   = len(boxes)
		indices []int
		answer  = make([]int, total)
	)
	for i, b := range boxes {
		if b == '1' {
			indices = append(indices, i)
		}
	}
	for i := 0; i < total; i++ {
		var t int
		for _, index := range indices {
			t += dis(i, index)
		}
		answer[i] = t
	}
	return answer
}

// O(N)
func minOperations(boxes string) []int {
	var (
		total = len(boxes)
		// 前缀和的前缀和
		preOrderSum  = make([]int, total)
		postOrderSum = make([]int, total)
		answer       = make([]int, total)
	)
	count, sum := 0, 0
	// 正序
	for i := 0; i < total; i++ {
		if boxes[i] == '1' {
			count++
		}
		sum += count
		preOrderSum[i] = sum
	}
	// 逆序
	count, sum = 0, 0
	for i := total - 1; i >= 0; i-- {
		if boxes[i] == '1' {
			count++
		}
		sum += count
		postOrderSum[i] = sum
	}
	for i := 0; i < total; i++ {
		sum := 0
		if i > 0 {
			sum += preOrderSum[i-1]
		}
		if i < total-1 {
			sum += postOrderSum[i+1]
		}
		answer[i] = sum
	}
	return answer
}

// O(N), 空间复杂度常数优化
func minOperationsLess(boxes string) []int {
	var (
		total  = len(boxes)
		answer = make([]int, total)
	)
	count, sum := 0, 0
	// 正序构建前缀和的前缀和
	for i := 0; i < total; i++ {
		if boxes[i] == '1' {
			count++
		}
		sum += count
		if i < total-1 {
			answer[i+1] += sum
		}
	}
	// 逆序构建前缀和的前缀和
	count, sum = 0, 0
	for i := total - 1; i >= 0; i-- {
		if boxes[i] == '1' {
			count++
		}
		sum += count
		if i > 0 {
			answer[i-1] += sum
		}
	}
	return answer
}

func main() {
	fmt.Println(minOperationsLess("001011"))
}
