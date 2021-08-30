package main

import (
	"github.com/stdioa/leetcode/linkedlist"
)

type Node = linkedlist.ListNode

func insert(aNode *Node, x int) *Node {
	var newNode = &Node{Val: x}
	if aNode == nil {
		newNode.Next = newNode
		return newNode
	}

	// 找到起始点
	var (
		startNode = aNode.Next
		p         = startNode.Next
	)
	for p != startNode && p.Next.Val >= p.Val {
		p = p.Next
	}
	// 两种情况跳出第二段查找：
	// 1. p 是一个处处相等的环
	// 2. p 处在在整个环的最大值，但还是比 x 要小
	if p.Val > x {
		for p.Next.Val <= x {
			p = p.Next
		}
	}
	newNode.Next = p.Next
	p.Next = newNode
	return aNode
}
