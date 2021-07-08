/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 大体思路是这样，写的有点啰嗦
func search(node, p, q *TreeNode) (result *TreeNode, hasP, hasQ bool) {
    hasP = (node == p
        hasP = true
    }
    if node == q {
        hasQ = true    
    }
    if hasP && hasQ {
        return node, hasP, hasQ
    }
    
    nodes := []*TreeNode{node.Left, node.Right}
    for _, n := range nodes {
        if n == nil {
            continue
        }
        r, pp, qq := search(n, p, q)
        result = r
        hasP = hasP || pp
        hasQ = hasQ || qq
        if result != nil {
            return result, hasP, hasQ
        } else if hasP && hasQ {
            return node, hasP, hasQ
        }
    }
    return nil, hasP, hasQ
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    result, _, _ := search(root, p, q)
    return result
}
