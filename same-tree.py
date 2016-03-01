# coding: utf-8

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    # @param {TreeNode} p
    # @param {TreeNode} q
    # @return {boolean}
    def isSameTree(self, p, q):
        def isParted(p, q):
            if p.val != q.val:                      # judge the value of current nodes
                return False

            if p.left and q.left:                   # both nodes have the left child
                ans = isParted(p.left, q.left)
                if not ans:
                    return False
            elif p.left != q.left:                  # a node have the left child, but the other doesn't
                return False
                
            if p.right and q.right:
                ans = isParted(p.right, q.right)    # judge the right child
                if not ans:
                    return False
            elif p.right != q.right:
                return False

            return True
        
        if p and q:
            return isParted(p, q)
        if p == q:
            return True
        return False
                