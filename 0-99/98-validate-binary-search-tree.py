# coding: utf-8

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

from made_bitree import TreeNode

class Solution(object):
    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if not root:
            return True
            
        def isValidNode(node):
            if not (node.left or node.right):
                return node.val, node.val, True

            max_ = min_ = node.val

            if node.left:
                _, max_, valid = isValidNode(node.left)
                if not valid:
                    return None, None, False
                elif max_ >= node.val:
                    print(max_, node.val)
                    return None, None, False

            if node.right:
                min_, _, valid = isValidNode(node.right)
                if not valid:
                    return None, None, False
                elif min_ <= node.val:
                    print(min_, node.val)
                    return None, None, False
            
            max_ = max(max_, node.val)
            min_ = min(min_, node.val)

            return min_, max_, True
            
        return isValidNode(root)[2]

tree = TreeNode("{10,5,15,#,#,6,20}")
# tree = TreeNode("{5, 14, 1}")
obj = Solution()
print(obj.isValidBST(tree))
