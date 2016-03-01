# coding: utf-8

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    # @param {TreeNode} root
    # @return {boolean}
    def isSymmetric(self, root):
        def isSymNode(left, right):
            if left.val != right.val:
                return False

            if left.left and right.right:
                ans = isSymNode(left.left, right.right)
                if not ans:
                    return False
            elif left.left != right.right:
                return False

            if left.right and right.left:
                ans = isSymNode(left.right, right.left)
                if not ans:
                    return False
            elif left.right != right.left:
                return False

            return True

        if not root:
            return True

        if root.left and root.right:
            return isSymNode(root.left, root.right)
        elif root.left != root.right:
            return False

        return True