package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// Get list length
	var (
		length int
		node   = head
	)
	for node != nil {
		length++
		node = node.Next
	}
	k = k % length
	if k == 0 {
		return head
	}

	var (
		back  = head
		front = head
	)
	for i := 1; i <= k; i++ {
		front = front.Next
	}
	for front.Next != nil {
		front = front.Next
		back = back.Next
	}
	newHead := back.Next
	back.Next = nil
	front.Next = head
	return newHead
}
