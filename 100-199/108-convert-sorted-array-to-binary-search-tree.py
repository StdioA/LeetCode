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
        def makePart(numl):
            if not numl:
                return None

            l = len(numl)/2
            pivot = numl[l]
            root = TreeNode(pivot)
            
            lower = numl[:l]
            upper = numl[l+1:]
            if lower:
                root.left = makePart(lower)
            if upper:
                root.right = makePart(upper)

            return root
        
        return makePart(nums)
        