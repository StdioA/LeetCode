package main

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

func countSmaller(nums []int) []int {
	// 构建树状数组
	var (
		length   = len(nums)
		numArray = NewNumArray(20001)
		results  []int
	)
	for i := length - 1; i >= 0; i-- {
		// 倒序遍历并查询，注意如果要查更小的数的话，num 需要减一
		num := nums[i] + 10000 + 1
		result := numArray.GetSum(num - 1)
		results = append(results, result)
		numArray.Add(num, 1)
	}

	// 翻转结果
	for i := 0; i < len(results)/2; i++ {
		results[i], results[length-i-1] = results[length-i-1], results[i]
	}
	return results
}
