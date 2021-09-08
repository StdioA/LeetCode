package main

import (
	"github.com/stdioa/leetcode/linkedlist"
)

type ListNode = linkedlist.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var (
		p        = head
		distinct = true
		newHead  = &ListNode{}
		newp     = newHead
	)
	for p.Next != nil {
		if p.Next.Val == p.Val {
			distinct = false
		} else {
			if distinct {
				newp.Next = p
				newp = p
			}
			distinct = true
		}
		p = p.Next
	}
	if distinct {
		newp.Next = p
		newp = p
	}
	newp.Next = nil
	return newHead.Next
}
