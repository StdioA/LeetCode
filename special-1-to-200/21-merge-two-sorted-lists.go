package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 归并排序链表版
	var (
		p1, p2  = list1, list2
		head    = &ListNode{}
		mergedP = head
	)
	for p1 != nil || p2 != nil { // 这个地方不要写错，!(p1 == nil && p2 == nil) 可能会更直观
		var minNode = p1
		if p2 != nil {
			if minNode == nil || p2.Val < minNode.Val {
				minNode = p2
			}
		}
		if minNode == p1 {
			p1 = p1.Next
		} else if minNode == p2 {
			p2 = p2.Next
		}
		mergedP.Next = minNode
		mergedP = minNode
	}
	return head.Next
}
