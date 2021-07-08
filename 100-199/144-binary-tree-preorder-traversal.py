# coding: utf-8

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

from made_bitree import TreeNode

class Solution:
    # @param {TreeNode} root
    # @return {integer[]}
    def preorderTraversal(self, root):
        if not root:
            return []
        queue = [root]
        ans = []
        while queue:
            node = queue.pop()
            # print(node.val)
            ans.append(node.val)
            if node.right:
                queue.append(node.right)
            if node.left:
                queue.append(node.left)
        return ans

tree = TreeNode('{1,4,3,2}')
print(Solution().preorderTraversal(tree))
