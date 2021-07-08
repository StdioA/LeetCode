package main

import (
	"fmt"

	"github.com/stdioa/leetcode/linkedlist"
)

type ListNode = linkedlist.ListNode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if m == n {
		return head
	}
	// Find head for reversing
	var phead, rhead *ListNode
	rhead = head
	for i := 1; i < m; i++ {
		phead, rhead = rhead, rhead.Next
	}
	// Reverse
	var prev, cur, next *ListNode
	prev, cur = rhead, rhead.Next
	for i := 1; i <= (n - m); i++ {
		next = cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	// Deal with links around the reversed list
	if phead != nil {
		phead.Next = prev
	} else {
		head = prev
	}
	rhead.Next = cur
	return head
}

func main() {
	list := linkedlist.New([]int{1, 2, 3, 4, 5})
	fmt.Println(reverseBetween(list, 2, 4))
}
