# coding: utf-8

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    # @param {TreeNode} root
    # @return {string[]}
    def binaryTreePaths(self, root):
        def find(node):
            if not node:
                return []
            
            paths = []
            left = find(node.left)
            right = find(node.right)
            vals = str(node.val)
            
            if (not left) and (not right):
                return [str(node.val)]
            else:
                left = ["->".join([vals, x]) for x in left]
                right = ["->".join([vals, x]) for x in right]
                return left+right
            
        return find(root)
