package main

import (
	"fmt"

	"github.com/stdioa/leetcode/linkedlist"
)

type ListNode = linkedlist.ListNode

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var (
		lessHead    = &ListNode{}
		gteHead     = &ListNode{}
		lessp, gtep = lessHead, gteHead
		p           = head
	)

	for p != nil {
		if p.Val < x {
			lessp.Next = p
			lessp = p
		} else {
			gtep.Next = p
			gtep = p
		}
		fmt.Println(p)
		p = p.Next
	}
	gtep.Next = nil

	var resHead = lessHead.Next
	if lessp == lessHead {
		resHead = gteHead.Next
	} else {
		lessp.Next = gteHead.Next
	}
	return resHead
}

func main() {
	list := linkedlist.New([]int{1, 4, 3, 2, 5, 2})
	fmt.Println(partition(list, 3))
}
