package main

import (
	"github.com/stdioa/leetcode/bitree"
	"github.com/stdioa/leetcode/linkedlist"
)

type (
	TreeNode = bitree.TreeNode
	ListNode = linkedlist.ListNode
)

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var (
		pre        *ListNode
		slow, fast = head, head
	)
	// 快慢指针求中点
	for fast.Next != nil && fast.Next.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	var node = &TreeNode{Val: slow.Val}
	// 切断左半边链表
	if pre != nil {
		pre.Next = nil
		node.Left = sortedListToBST(head)
	}

	// 切断右半边链表
	right := slow.Next
	slow.Next = nil
	node.Right = sortedListToBST(right)
	return node
}
