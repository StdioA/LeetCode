package main

type Solution struct {
	indices       map[int][]int
	traverseIndex map[int]int
}

func Constructor(nums []int) Solution {
	var sol = Solution{
		indices:       make(map[int][]int),
		traverseIndex: make(map[int]int),
	}
	for i, num := range nums {
		sol.indices[num] = append(sol.indices[num], i)
	}
	return sol
}

func (this *Solution) Pick(target int) int {
	var (
		index  = this.traverseIndex[target]
		result = this.indices[target][index]
	)
	this.traverseIndex[target] = (index + 1) % len(this.indices[target])
	return result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
