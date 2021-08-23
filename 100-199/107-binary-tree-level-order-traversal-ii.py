class Solution:
    # @param root, a tree node
    # @return a list of lists of integers
    def levelOrderBottom(self, root):
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
        tra.reverse()
        return tra
