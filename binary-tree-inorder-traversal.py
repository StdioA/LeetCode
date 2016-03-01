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
    def inorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if not root:
            return []

        stack = [(root, 0)]                                 # 0表示该节点未开始遍历，1表示左子树遍历中，2表示右
        ans = []

        while stack:
            node, status = stack.pop()
            if not (node.left or node.right):               # 当前节点为叶子节点
                ans.append(node.val)
            else:
                if status == 0:                             # 下一步应遍历左子树
                    stack.append((node, 1))
                    if node.left:
                        stack.append((node.left, 0))

                elif status == 1:                           # 下一步应输出当前节点并遍历右子树
                    stack.append((node, 2))
                    ans.append(node.val)
                    if node.right:
                        stack.append((node.right, 0))

                elif status == 2:                           # 什么都不做
                    pass
            
        return ans

tree = TreeNode('{1,2,3,4,5,6,7}')
print(Solution().inorderTraversal(tree))
