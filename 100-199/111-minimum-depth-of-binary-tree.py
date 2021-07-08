# coding: utf-8

# Definition for a  binary tree node
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

from made_bitree import TreeNode

class Solution:
    # @param root, a tree node
    # @return an integer
    def minDepth(self, root):
        # the code comes from the problem "Binary tree level order traversal"
        pre = []
        cur = [root]
        tra = []
        level = 0
        if root:
            while cur:
                pre, cur = cur[:], []
                level += 1
                for node in pre:
                    if not (node.left or node.right):
                        return level
                    if node.left:
                        cur.append(node.left)
                    if node.right:
                        cur.append(node.right)
        else:
            return 0

# root = TreeNode("{3,9,20,#,#,15,7}")
root = TreeNode("{1,2}")

a = Solution()
print a.minDepth(root)