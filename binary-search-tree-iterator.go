/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	root, cur *TreeNode
	stack     []*TreeNode
	p         *TreeNode
	next      int
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{
		root:  root,
		stack: []*TreeNode{root},
		p:     root,
	}
	iter.next = iter.flush()
	return iter
}

func (this *BSTIterator) flush() int {
	for this.p != nil || len(this.stack) > 0 {
		for this.p != nil {
			this.stack = append(this.stack, this.p)
			this.p = this.p.Left
		}
		if len(this.stack) > 0 {
			this.p = this.stack[len(this.stack)-1]
			val := this.p.Val
			this.stack = this.stack[:len(this.stack)-1]
			this.p = this.p.Right
			return val
		}
	}
	return 0
}

func (this *BSTIterator) Next() int {
	next := this.next
	this.next = this.flush()
	return next
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */