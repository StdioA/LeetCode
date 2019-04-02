package main

import (
	"fmt"

	"github.com/stdioa/leetcode/linkedlist"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

func main() {
	list := linkedlist.NewWithCycle([]int{3, 2, 0, 4}, 1)
	node := detectCycle(list)
	fmt.Println(node.Val) // 2
}
