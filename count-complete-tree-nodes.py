# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

from made_bitree import TreeNode


class Solution:
    def depth(self, root: TreeNode) -> int:
        if root is None:
            return 0
        depth = 1
        while root.left is not None:
            depth += 1
            root = root.left
        return depth

    def countNodes(self, root: TreeNode) -> int:
        if root is None:
            return 0

        depth = self.depth(root)
        if self.depth(root.right) == depth - 1:
            return (1 << (depth - 1)) + self.countNodes(root.right)
        else:
            return (1 << (depth - 2)) + self.countNodes(root.left)


a = Solution()
tree = TreeNode("{1,2,3,4,5,6}")
print(a.countNodes(tree))
