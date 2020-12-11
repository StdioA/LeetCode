package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		p1, p2 = l1, l2
		root   = &ListNode{}
		pr     = root
	)
	for (p1 != nil) && (p2 != nil) {
		if p1.Val < p2.Val {
			pr.Next = p1
			p1 = p1.Next
		} else {
			pr.Next = p2
			p2 = p2.Next
		}
		pr = pr.Next
	}
	if p1 != nil {
		pr.Next = p1
	} else if p2 != nil {
		pr.Next = p2
	}

	return root.Next
}
