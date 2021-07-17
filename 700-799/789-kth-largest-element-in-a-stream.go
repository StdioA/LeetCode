package main

type KthLargest struct {
	heap []int // 小顶堆
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	// build heap
	this := KthLargest{
		k:    k,
		heap: []int{},
	}
	for _, num := range nums {
		this.Add(num)
	}
	return this
}

func (this *KthLargest) down(i int) {
	var l = len(this.heap)
	for i < l {
		maxi := i
		il, ir := i*2, i*2+1
		if il < l && this.heap[il] < this.heap[maxi] {
			maxi = il
		}
		if ir < l && this.heap[ir] < this.heap[maxi] {
			maxi = ir
		}
		if maxi == i {
			return
		}
		this.heap[i], this.heap[maxi] = this.heap[maxi], this.heap[i]
		i = maxi
	}
}

func (this *KthLargest) up(i int) int {
	for i > 0 {
		father := i / 2
		if this.heap[i] <= this.heap[father] {
			this.heap[i], this.heap[father] = this.heap[father], this.heap[i]
			this.down(i)
			i = father
		} else {
			break
		}
	}
	return i
}

func (this *KthLargest) Add(val int) int {
	if len(this.heap) < this.k {
		this.heap = append(this.heap, val)
		this.down(this.up(len(this.heap) - 1))
		return this.heap[0]
	}

	if val > this.heap[0] {
		this.heap[0] = val
		this.down(0)
	}
	return this.heap[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
