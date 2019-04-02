package linkedlist

import (
	"fmt"
	"strconv"
	"strings"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// New Make a linked list from an array
func New(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	node := head
	for _, num := range nums[1:] {
		node.Next = &ListNode{Val: num}
		node = node.Next
	}
	return head
}

// NewWithCycle Make a linked list from an array, and generate a cycle
func NewWithCycle(nums []int, cycle int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	nodes := []*ListNode{head}
	node := head
	for _, num := range nums[1:] {
		node.Next = &ListNode{Val: num}
		node = node.Next
		nodes = append(nodes, node)
	}
	node.Next = nodes[cycle]
	return head
}

// AString func
func (root *ListNode) String() string {
	result := make([]string, 0)
	nodes := make(map[*ListNode]int)
	head := root
	index := 0
	cycle := -1
	for head != nil {
		i, ext := nodes[head]
		if ext {
			cycle = i
			break
		}
		nodes[head] = index
		result = append(result, strconv.Itoa(head.Val))
		head = head.Next
		index++
	}
	if cycle == -1 {
		return fmt.Sprintf("[%s]", strings.Join(result, ","))
	}
	return fmt.Sprintf("[%s] -> %d", strings.Join(result, ","), cycle)
}
