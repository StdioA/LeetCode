package main

import (
	"github.com/stdioa/leetcode/linkedlist"
)

type ListNode = linkedlist.ListNode

func reorderList(head *ListNode) {
	// 快慢指针求出前一半和后一半
	var slow, fast = head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = slow.Next
	slow.Next = nil
	// 将后一半翻转
	var reverse *ListNode
	for fast != nil {
		nextp := fast.Next
		fast.Next = reverse
		reverse = fast
		fast = nextp
	}
	// 交叉合并两个链表
	var p, r = head, reverse
	for p != nil {
		var (
			nextp = p.Next
			nextr *ListNode
		)
		p.Next = r
		if r != nil {
			nextr = r.Next
			r.Next = nextp
		}
		p, r = nextp, nextr
	}
}
