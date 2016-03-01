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
    # @param {integer} sum
    # @return {boolean}
    def hasPathSum(self, root, sum):
        def check(root, sum_):
            next = sum_-root.val
            if next == 0 and not (root.left or root.right):
                return True
            else:
                if root.left:
                    ans = check(root.left, next)
                    if ans:
                        return True
                if root.right:
                    ans = check(root.right, next)
                    if ans:
                        return True
            return False

        if not root:
            return False
        else:
            return check(root, sum)

# a = TreeNode('{5,4,8,11,#,13,4,7,2,#,#,#,1}')
a = TreeNode('{-2,#,-3}')
print Solution().hasPathSum(a,-5)