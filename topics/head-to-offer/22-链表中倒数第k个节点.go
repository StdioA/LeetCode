package main

import (
	"github.com/stdioa/leetcode/linkedlist"
)

type ListNode = linkedlist.ListNode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var (
		prev, next = head, head
	)
	for i := 0; i < k; i++ {
		next = next.Next
	}
	for next != nil {
		prev, next = prev.Next, next.Next
	}
	return prev
}
