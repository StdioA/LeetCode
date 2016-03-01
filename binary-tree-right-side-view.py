# coding: utf-8

# Definition for a  binary tree node
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

from made_bitree import TreeNode

class Solution:
    def rightSideView(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
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
        ans = [n[-1] for n in tra]
        return ans

root = TreeNode("{3,9,20,#,#,15,7}")

a = Solution()
print(a.rightSideView(root))
