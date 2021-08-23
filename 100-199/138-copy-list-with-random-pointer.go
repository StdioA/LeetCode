package main

type Node struct {
	Val          int
	Next, Random *Node
}

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

func copyRandomListWithConstantSpace(head *Node) *Node {
	if head == nil {
		return head
	}
	var p = head
	// Insert copied node right after original node
	for p != nil {
		copyP := &Node{
			Val:  p.Val,
			Next: p.Next,
		}
		p.Next = copyP
		p = copyP.Next
	}
	// Fill random pointer
	p = head
	for p != nil {
		newP := p.Next
		if p.Random != nil {
			newP.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	// Split the nodes
	var (
		newHead = head.Next
		newP    = newHead
	)
	p = head
	for p.Next.Next != nil {
		nextP := p.Next.Next

		newP.Next = nextP.Next
		newP = nextP.Next
		p.Next = nextP
		p = nextP
	}
	p.Next = nil

	return newHead
}
