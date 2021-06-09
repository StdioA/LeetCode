package main

import "fmt"

// 树状数组
type NumArray struct {
	num []int
	sum []int
}

func lowbit(n int) int {
	return n & (-n)
}

func Constructor(nums []int) NumArray {
	var array = NumArray{
		num: append([]int{0}, nums...),
		sum: make([]int, len(nums)+1),
	}
	for i := 1; i < len(array.num); i++ {
		array.sum[i] = array.num[i]
		for j := i - 1; j > i-lowbit(i); j -= lowbit(j) {
			array.sum[i] += array.sum[j]
		}
	}
	return array
}

func (this *NumArray) Update(index int, val int) {
	index++
	delta := val - this.num[index]
	this.num[index] = val

	for index < len(this.num) {
		this.sum[index] += delta
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
	ans := this.getSum(right + 1)
	if left > 0 {
		ans -= this.getSum(left)
	}
	return ans
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

func main() {
	array := Constructor([]int{7, 9, 7, 2, 0})
	fmt.Println(array.SumRange(0, 2))
	array.Update(1, 2)
	fmt.Println(array.SumRange(0, 2))
}
