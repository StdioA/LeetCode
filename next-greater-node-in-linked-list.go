import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return make([]int, 0)
	}

	// get length
	var length, index int
	node := head
	for head != nil {
		length++
		head = head.Next
	}

	result := make([]int, length)
	nodeStack := make([]*ListNode, 0)
	indexStack := make([]int, 0)
	for node != nil {
		// pop from stack
		length := len(nodeStack)
		for length > 0 && nodeStack[length-1].Val < node.Val {
			fmt.Println(indexStack)
			i := indexStack[length-1]
			result[i] = node.Val

			nodeStack = nodeStack[:length-1]
			indexStack = indexStack[:length-1]
			length--
		}

		indexStack = append(indexStack, index)
		nodeStack = append(nodeStack, node)
		index++
		node = node.Next
	}
	return result
}
