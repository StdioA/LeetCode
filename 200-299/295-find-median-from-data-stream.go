package main

import "fmt"

type heap struct {
	cmp  func(i, j int) bool
	nums []int
}

func (h *heap) Top() int {
	return h.nums[0]
}

func (h *heap) adjustDown(i int) {
	var l = len(h.nums)
	for i < l {
		maxi := i
		left, right := i*2+1, i*2+2
		if left < l && h.cmp(h.nums[left], h.nums[maxi]) {
			maxi = left
		}
		if right < l && h.cmp(h.nums[right], h.nums[maxi]) {
			maxi = right
		}
		if maxi == i {
			return
		}
		h.nums[maxi], h.nums[i] = h.nums[i], h.nums[maxi]
		i = maxi
	}
}

func (h *heap) Push(n int) {
	h.nums = append(h.nums, 0)
	l := len(h.nums)
	copy(h.nums[1:], h.nums[:l-1])
	h.nums[0] = n
	h.adjustDown(0)
}

func (h *heap) Pop() int {
	top := h.nums[0]
	l := len(h.nums)
	h.nums[0], h.nums[l-1] = h.nums[l-1], h.nums[0]
	h.nums = h.nums[:l-1]
	h.adjustDown(0)
	return top
}

func (h *heap) Len() int {
	return len(h.nums)
}

func newHeap(max bool) *heap {
	h := &heap{
		cmp: func(i, j int) bool {
			return i < j
		},
	}
	if max {
		h.cmp = func(i, j int) bool {
			return i > j
		}
	}
	return h
}

type MedianFinder struct {
	lowHeap, highHeap *heap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		lowHeap:  newHeap(true),
		highHeap: newHeap(false),
	}
}

func (this *MedianFinder) balance() {
	for this.lowHeap.Len()-this.highHeap.Len() > 1 {
		this.highHeap.Push(this.lowHeap.Pop())
	}

	for this.highHeap.Len()-this.lowHeap.Len() > 1 {
		this.lowHeap.Push(this.highHeap.Pop())
	}
}

func (this *MedianFinder) AddNum(num int) {
	// defer this.print()
	if float64(num) > this.FindMedian() {
		this.highHeap.Push(num)
	} else {
		this.lowHeap.Push(num)
	}
	this.balance()
}

func (this *MedianFinder) print() {
	fmt.Printf("%v %v\n", this.lowHeap, this.highHeap)
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lowHeap.Len() == 0 && this.highHeap.Len() == 0 {
		return -100
	} else if this.lowHeap.Len() == this.highHeap.Len() {
		return float64(this.lowHeap.Top()+this.highHeap.Top()) / 2.00

	} else if this.lowHeap.Len() > this.highHeap.Len() {
		return float64(this.lowHeap.Top())

	} else {
		return float64(this.highHeap.Top())
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	h := Constructor()
	h.AddNum(2)
	h.AddNum(1)
	fmt.Println(h.FindMedian(), 1.5)
	h.AddNum(3)
	fmt.Println(h.FindMedian(), 2)
}
