func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		la, lb int
		pa, pb = headA, headB
	)
	for pa != nil {
		la++
		pa = pa.Next
	}
	for pb != nil {
		lb++
		pb = pb.Next
	}
	pa, pb = headA, headB
	if la > lb {
		for i := 0; i < la-lb; i++ {
			pa = pa.Next
		}
	} else {
		for i := 0; i < lb-la; i++ {
			pb = pb.Next
		}
	}
	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa, pb = pa.Next, pb.Next
	}

	return nil
}
