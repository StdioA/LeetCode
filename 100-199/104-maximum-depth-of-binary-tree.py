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
    def maxDepth(self, root):
        """ Doing the level-order-traverse is FAST"""
        # the code comes from the problem "Binary tree level order traversal"
        pre = []
        cur = [root]
        tra = []
        if root:
            while cur:
                tra.append([])
                pre, cur = cur[:], []
                for node in pre:
                    tra[-1].append(node.val)
                    if node.left:
                        cur.append(node.left)
                    if node.right:
                        cur.append(node.right)
        return len(tra)

    def maxDepth_TLE(self, root):
        """ The recursion with factitial stack is really SLOW! """
        if not root:
            return 0
        max_ = level = 1
        pointer = root
        stack = [root]

        lvisited, rvisited = [], []

        while stack:
            pointer = stack[-1]
            if (pointer not in lvisited) and pointer.left:
                lvisited.append(pointer)
                stack.append(pointer.left)
                level += 1
            elif (pointer not in rvisited) and pointer.right:
                rvisited.append(pointer)
                stack.append(pointer.right)
                level += 1
            else:
                max_ = max(max_, level)
                stack.pop()
                level -= 1
        return max_

root = TreeNode("{3,9,20,#,#,15,7}")

a = Solution()
print a.maxDepth(root)