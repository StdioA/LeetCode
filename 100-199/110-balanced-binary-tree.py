# coding: utf-8

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isBalanced(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        def check(node):
            if not node:
                return 0
            left = check(node.left)
            right = check(node.right)
            
            if (left is None) or (right is None):
                return None
            elif abs(left-right) > 1:
                return None
            else:
                return max(left, right)+1
            
        ans = check(root)
        return (ans is not None)

