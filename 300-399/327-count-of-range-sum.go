package main

import (
	"fmt"
	"sort"
)

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

func (this *NumArray) getSum(index int) int {
	var ans int
	for index > 0 {
		ans += this.sum[index]
		index -= lowbit(index)
	}
	return ans
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.getSum(right) - this.getSum(left-1)
}

// 设 sum 为前缀和，针对任意 i <= j
// lower <= sum[j] - sum[i-1] <= upper
// lower + sum[i-1] <= sum[j] <= upper + sum[i-1]
// 针对每个 j >= i，查询该范围内 sum[j] 出现的次数之和
// 问题变成了求某个范围内的前缀和出现的次数
// “单点更新，范围查找”

// PS: 逆序遍历前缀和时是上面那条，正序遍历时用下面这条
// sum[j] - upper <= sum[i-1] <= sum[j] - lower
// Credit to: https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0327.Count-of-Range-Sum/
func countRangeSum(nums []int, lower int, upper int) int {
	var (
		length    = len(nums)
		prefixSum = make([]int, length+1)
		allNums   = make([]int, 0, length*3+3)
	)
	allNums = append(allNums, 0, lower, upper)
	for i, n := range nums {
		prefixSum[i+1] = prefixSum[i] + n
		allNums = append(allNums, prefixSum[i+1], prefixSum[i+1]+lower, prefixSum[i+1]+upper)
	}

	// 离散化（说实话我觉得叫密集化更合适_(:з」∠)_）
	sort.Ints(allNums)
	var (
		indexMap = map[int]int{
			allNums[0]: 1,
		}
		k = 1
	)
	for i := 1; i < len(allNums); i++ {
		if allNums[i] != allNums[i-1] {
			k++
			indexMap[allNums[i]] = k
		}
	}

	// 树状数组
	var (
		numArray = NewNumArray(k)
		total    int
	)
	// 倒序插入
	for i := length; i > 0; i-- {
		sum := prefixSum[i]
		lowerBound, upperBound := prefixSum[i-1]+lower, prefixSum[i-1]+upper
		numArray.Add(indexMap[sum], 1)
		total += numArray.SumRange(indexMap[lowerBound], indexMap[upperBound])
	}
	return total
}

func main() {
	arr := []int{-2, 5, -1}
	fmt.Println(countRangeSum(arr, -2, 2))

	arr = []int{-2, -1, -2, -1, -3, -2}
	fmt.Println(countRangeSum(arr, 1, 6))
}
