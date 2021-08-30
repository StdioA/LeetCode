package main

func reversePrint(head *ListNode) []int {
	var (
		result []int
		p      = head
		l      int
	)
	for p != nil {
		result = append(result, p.Val)
		p = p.Next
		l++
	}
	for i := 0; i < l/2; i++ {
		result[i], result[l-i-1] = result[l-i-1], result[i]
	}
	return result
}
