package main

func mergeKLists(lists []*ListNode) *ListNode {
	// 在归并两个的流程上改一下就好了
	var (
		ps      = make([]*ListNode, len(lists))
		head    = &ListNode{}
		mergedP = head
	)
	for i, list := range lists {
		ps[i] = list
	}
	for {
		var (
			minNode  *ListNode
			minIndex = -1
		)
		for i, node := range ps {
			if node == nil {
				continue
			}
			if minNode == nil || node.Val < minNode.Val {
				minNode = node
				minIndex = i
			}
		}
		if minIndex >= 0 {
			mergedP.Next = minNode
			mergedP = minNode
			ps[minIndex] = ps[minIndex].Next
		} else {
			break
		}
	}
	return head.Next
}
