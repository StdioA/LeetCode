# coding: utf-8

from made_bitree import TreeNode

class Solution(object):
    def invertTree(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        if not root:
            return root
            
        def invertNode(node):
            node.left, node.right = node.right, node.left
            if node.left:
                invertNode(node.left)
            if node.right:
                invertNode(node.right)
                
        invertNode(root)
        
        return root

obj = Solution()
tree = TreeNode("{4,2,7,1,3,6,9}")
print(str(obj.invertTree(tree)))
