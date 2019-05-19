package main

import (
	"fmt"
	"github.com/stdioa/leetcode/linkedlist"
)

type ListNode=linkedlist.ListNode


func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pp *ListNode = nil
	prev, next := head, head.Next
	head = next
	for prev != nil && next != nil {
		prev.Next, next.Next = next.Next, prev
		if pp != nil {
			pp.Next = next
		}
		
		pp, prev = prev, prev.Next
		if prev != nil {
			next = prev.Next
		}
	}
	return head
}

func main() {
	list := linkedlist.New([]int{1,2})
	fmt.Println(swapPairs(list))
}
