package main

import "github.com/stdioa/leetcode/linkedlist"

type ListNode = linkedlist.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 简单的竖式计算，需要注意可能有进位
	var (
		result  = new(ListNode)
		resultP = result
		p1, p2  = l1, l2
		carry   int
	)
	for p1 != nil || p2 != nil {
		var sum = carry
		if p1 != nil {
			sum += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			sum += p2.Val
			p2 = p2.Next
		}
		carry = sum / 10
		resultP.Next = &ListNode{Val: sum % 10}
		resultP = resultP.Next

	}
	if carry > 0 {
		resultP.Next = &ListNode{Val: carry}
		resultP = resultP.Next
	}
	return result.Next
}
