package main

import (
	"fmt"

	"github.com/stdioa/leetcode/linkedlist"
)

type ListNode = linkedlist.ListNode

func sortList(head *ListNode) *ListNode {
	// Divide
	var prev *ListNode
	l1, l2 := head, head
	for l2 != nil && l2.Next != nil {
		prev = l1
		l1, l2 = l1.Next, l2.Next
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if l1 == l2 {
		return head
	}
	prev.Next = nil

	// Sort
	l1, l2 = sortList(head), sortList(l1)

	// Merge
	result := &ListNode{}		// Virtual head
	resultHead := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			result.Next = l1
			l1 = l1.Next
		} else {
			result.Next = l2
			l2 = l2.Next
		}
		result = result.Next
	}
	for l1 != nil {
		result.Next = l1
		l1 = l1.Next
		result = result.Next
	}
	for l2 != nil {
		result.Next = l2
		l2 = l2.Next
		result = result.Next
	}
	return resultHead.Next
}

func main() {
	l := linkedlist.New([]int{-1, 5, 3, 4})	
	fmt.Println(sortList(l))
}
