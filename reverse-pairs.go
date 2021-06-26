package main

import "sort"

// 树状数组
type NumArray struct {
	sum []int
}

func lowbit(n int) int {
	return n & (-n)
}

func NewNumArray(amount int) *NumArray {
	var array = &NumArray{
		sum: make([]int, amount+1),
	}
	return array
}

func (this *NumArray) Add(index int, val int) {

	for index < len(this.sum) {
		this.sum[index] += val
		index += lowbit(index)
	}
}

func (this *NumArray) GetSum(index int) int {
	var ans int
	for index > 0 {
		ans += this.sum[index]
		index -= lowbit(index)
	}
	return ans
}

func reversePairs(nums []int) int {
	var (
		length       = len(nums)
		copyNums     = make([]int, length*2)
		k        int = 1
	)
	for i, num := range nums {
		copyNums[i*2], copyNums[i*2+1] = num, num*2
	}
	// 倒排所有数
	sort.Sort(sort.Reverse(sort.IntSlice(copyNums)))
	// 散列化
	indexMap := map[int]int{
		copyNums[0]: k,
	}
	for i := 1; i < len(copyNums); i++ {
		if copyNums[i] != copyNums[i-1] {
			k++
			indexMap[copyNums[i]] = k
		}
	}

	var (
		numArray = NewNumArray(k)
		total    int
	)

	for _, num := range nums {
		total += numArray.GetSum(indexMap[num*2] - 1)
		numArray.Add(indexMap[num], 1)
	}
	return total
}
