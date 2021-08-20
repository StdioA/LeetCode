package main

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var (
		newHead = &Node{Val: head.Val}
		nodeMap = map[*Node]*Node{
			head: newHead,
		}
		p, newp = head, newHead
	)
	for p.Next != nil {
		node := &Node{
			Val: p.Next.Val,
		}
		newp.Next = node
		p, newp = p.Next, newp.Next
		nodeMap[p] = newp
	}

	// Assign random value
	p, newp = head, newHead
	for p != nil {
		newp.Random = nodeMap[p.Random]
		p, newp = p.Next, newp.Next
	}
	return newHead
}
