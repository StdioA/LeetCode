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

func (this *NumArray) GetSum(index int) int {
	var ans int
	for index > 0 {
		ans += this.sum[index]
		index -= lowbit(index)
	}
	return ans
}

// 树状数组
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

func mergeSort(arr []int) ([]int, int) {
	var length = len(arr)
	if length <= 1 {
		return arr, 0
	}
	leftArr, leftSum := mergeSort(arr[:length/2])
	rightArr, rightSum := mergeSort(arr[length/2:])

	// Accumulate result
	var result = leftSum + rightSum
	for ri, li := 0, 0; ri < len(rightArr); ri++ {
		// 这段和归并很像，也是两个指针交替前进
		// 如果每次都要把 li 置为 0，这段的时间复杂度就变成 O(n^2) 了，但是在外面初始化就是 O(n)
		for li < len(leftArr) && leftArr[li] <= 2*rightArr[ri] {
			li++
		}
		result += len(leftArr) - li
	}

	// Merge arrays
	var (
		buffer = make([]int, 0)
		li, ri = 0, 0
	)
	for li < len(leftArr) && ri < len(rightArr) {
		ln, rn := leftArr[li], rightArr[ri]
		if ln < rn {
			buffer = append(buffer, ln)
			li++
		} else {
			buffer = append(buffer, rn)
			ri++
		}
	}
	if li < len(leftArr) {
		buffer = append(buffer, leftArr[li:]...)
	}
	if ri < len(rightArr) {
		buffer = append(buffer, rightArr[ri:]...)
	}
	copy(arr, buffer)
	return arr, result
}

// 归并排序
func reversePairs2(arr []int) int {
	_, result := mergeSort(arr)
	return result
}

func main() {
	arr := []int{2, 4, 3, 5, 1}
	fmt.Println(reversePairs(arr))
	fmt.Println(reversePairs2(arr))
	arr = []int{5, 4, 3, 2, 1}
	fmt.Println(reversePairs2(arr))
}
