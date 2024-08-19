package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		front, end = head, head
	)
	for i := 0; i < n; i++ {
		end = end.Next
	}
	// 特殊情况，需要移除第一个元素
	if end == nil {
		return front.Next
	}
	// front 和 end 相差 n
	// 当 end 在最后一个元素时，front 就指向要被移除元素的前一个元素
	for end.Next != nil {
		front, end = front.Next, end.Next
	}
	front.Next = front.Next.Next
	return head
}
