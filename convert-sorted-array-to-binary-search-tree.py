# coding: utf-8

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    # @param {integer[]} nums
    # @return {TreeNode}
    def sortedArrayToBST(self, nums):
        if not nums:
            return
        def makePart(root, numl):
            l = len(numl)/2
            pivot = numl[l]
            root.val = pivot
            
            lower = numl[:l]
            upper = numl[l+1:]
            if lower:
                root.left = TreeNode(0)
                makePart(root.left, lower)
            if upper:
                root.right = TreeNode(0)
                makePart(root.right, upper)
        
        root = TreeNode(0)        
        makePart(root,nums)
        return root
        