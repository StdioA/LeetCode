package main

import (
	"fmt"

	"github.com/stdioa/leetcode/linkedlist"
)

type ListNode = linkedlist.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var (
		root = &ListNode{}
		p    = root
	)
	for {
		var (
			minNode  *ListNode
			minIndex int = -1
		)
		for i, node := range lists {
			if node == nil {
				continue
			}
			// Select minimal node
			if minNode == nil || node.Val < minNode.Val {
				minNode = node
				minIndex = i
			}
		}
		if minIndex == -1 {
			break
		}
		p.Next = minNode
		p = minNode
		lists[minIndex] = minNode.Next
	}
	return root.Next
}

func main() {
	lists := []*ListNode{
		linkedlist.New([]int{1, 4, 5}),
		linkedlist.New([]int{1, 3, 4}),
		linkedlist.New([]int{2, 6}),
	}
	node := mergeKLists(lists)
	fmt.Println(node)
}
