# coding: utf-8

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
from made_bitree import TreeNode


class Solution(object):
    def sumNumbers(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0
            
        def exe(node):
            total = 0

            if not(node.left or node.right):
                return node.val
            else:
                if node.left:
                    node.left.val += node.val*10
                    total += exe(node.left)
                if node.right:
                    node.right.val += node.val*10
                    total += exe(node.right)

                return total

        return exe(root)

a = Solution()
tree = TreeNode("{1,2,3}")
print a.sumNumbers(tree)
