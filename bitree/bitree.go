package bitree

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TreeNode Node for binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MakeBiTree Generate an binary tree from a string
func MakeBiTree(str string) (*TreeNode, error) {
	length := len(str)
	s, e := str[0], str[length-1]
	if !((s == '{' && e == '}') || (s == '[' && e == ']')) {
		return nil, errors.New("Bad format")
	}

	intfs := make([]interface{}, 0)
	for _, str := range strings.Split(str[1:length-1], ",") {
		str = strings.TrimSpace(str)
		if str == "nil" || str == "null" || str == "#" {
			intfs = append(intfs, nil)
		} else {
			val, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			intfs = append(intfs, val)
		}
	}
	return MakeBiTreeIntf(intfs), nil
}

func New(str string) (*TreeNode, error) {
	return MakeBiTree(str)
}

// MakeBiTreeIntf Generate an binary tree from an interface slice
func MakeBiTreeIntf(list []interface{}) *TreeNode {
	if len(list) == 0 {
		return nil
	}

	root := &TreeNode{Val: list[0].(int)}
	nodeQueue := []*TreeNode{root}
	for i, num := range list[1:] {
		if num != nil {
			currNode := nodeQueue[0]
			node := &TreeNode{Val: num.(int)}
			nodeQueue = append(nodeQueue, node)
			if i%2 == 0 { // operate left
				currNode.Left = node
			} else { // operate right
				currNode.Right = node
			}
		}
		if i%2 != 0 {
			nodeQueue = nodeQueue[1:]
		}
	}
	return root
}

func (root *TreeNode) String() string {
	layerIsEmpty := func(layer []*TreeNode) bool {
		for _, node := range layer {
			if node != nil {
				return false
			}
		}
		return true
	}

	layers := make([][]*TreeNode, 0)
	// Layer traverse tree while keeping nil pointers
	currentLayer := []*TreeNode{root}
	for !layerIsEmpty(currentLayer) {
		layers = append(layers, currentLayer)
		nextLayer := make([]*TreeNode, 0)
		for _, node := range currentLayer {
			if node != nil {
				nextLayer = append(nextLayer, node.Left, node.Right)
			}
		}
		currentLayer = nextLayer
	}
	// Combine layers
	result := make([]string, 0)
	for _, layer := range layers {
		for _, node := range layer {
			if node == nil {
				result = append(result, "nil")
			} else {
				result = append(result, strconv.Itoa(node.Val))
			}
		}
	}
	// Drop "nil"s at the tail
	length := len(result)
	i := length - 1
	for i > 0 && result[i] == "nil" {
		i--
	}
	result = result[:i+1]
	return fmt.Sprintf("{%s}", strings.Join(result, ","))
}

func preOrderT(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	result := []int{root.Val}
	result = append(result, preOrderT(root.Left)...)
	result = append(result, preOrderT(root.Right)...)
	return result
}

// PreOrderTraverse func
func PreOrderTraverse(root *TreeNode) []int {
	return preOrderT(root)
}
