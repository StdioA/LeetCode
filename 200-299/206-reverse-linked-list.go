package main

import (
	"fmt"

	"github.com/stdioa/leetcode/linkedlist"
)

type ListNode = linkedlist.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}

func main() {
	list := linkedlist.New([]int{1, 2, 3, 4, 5})
	fmt.Println(reverseList(list))
}
