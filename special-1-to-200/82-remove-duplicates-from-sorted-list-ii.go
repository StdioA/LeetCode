package main

func deleteDuplicates(head *ListNode) *ListNode {
	var (
		newHead = new(ListNode)
		newp    = newHead
		p       = head
	)
	newHead.Next = head
	for p != nil {
		if p.Next != nil && p.Val == p.Next.Val {
			// 若发现重复节点，则跳过后面所有相同值的节点
			for p.Next != nil && p.Val == p.Next.Val {
				p = p.Next
			}
			// 然后把 newp 和第一个不同的节点接起来
			newp.Next = p.Next
			p = p.Next
		} else {
			// 后面节点和当前节点不同，说明当前节点没有重复
			// 链表已经接上了，直接向后遍历就行
			newp = newp.Next
			p = p.Next
		}
	}
	return newHead.Next
}
