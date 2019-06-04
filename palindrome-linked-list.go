package main

import (
	"fmt"

	"github.com/stdioa/leetcode/linkedlist"
)

type ListNode = linkedlist.ListNode

func isPalindrome(head *ListNode) bool {
    // Calculate length
    length := 0
    var prev, cur, next *ListNode
    cur = head
    for cur != nil {
        cur = cur.Next
        length++
    }
    // Reverse previous half of list
    cur = head
    for i := 0; i < length / 2; i++ {
        next = cur.Next
        cur.Next = prev
        prev, cur = cur, next
    }
    if length % 2 == 1 {
        cur = cur.Next
    }
    for cur != nil {
        if cur.Val != prev.Val {
            return false
        }
        cur, prev = cur.Next, prev.Next
    }
    return true
}


func main() {
	l1 := linkedlist.New([]int{1, 2, 3, 2, 1})
	fmt.Println(isPalindrome(l1))
}
