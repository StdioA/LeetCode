package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head   = &ListNode{}
		p      = head
		p1, p2 = l1, l2
		carry  int
	)
	for p1 != nil && p2 != nil {
		sum := p1.Val + p2.Val + carry
		carry = sum / 10
		node := &ListNode{Val: sum % 10}
		p.Next = node
		p = p.Next

		p1, p2 = p1.Next, p2.Next
	}

	var pre = p
	if p1 != nil {
		p.Next = p1
	} else if p2 != nil {
		p.Next = p2
	}
	p = p.Next
	// Traverse list left
	for p != nil && carry > 0 {
		p.Val += carry
		carry = p.Val / 10
		p.Val %= 10
		pre, p = p, p.Next
	}
	// Deal with the last carry
	if carry > 0 {
		pre.Next = &ListNode{Val: carry}
	}

	return head.Next
}
