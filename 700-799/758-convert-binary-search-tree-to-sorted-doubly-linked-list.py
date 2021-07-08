from made_bitree import TreeNode

class Solution:
    def genList(self, node):
        if node is None:
            return None, None

        lhead, ltail = self.genList(node.left)
        rhead, rtail = self.genList(node.right)
        
        if ltail is not None:
            ltail.right, node.left = node, ltail
        else:
            lhead = node
        if rhead is not None:
            rhead.left, node.right = node, rhead
        else:
            rtail = node

        return lhead, rtail

    def treeToDoublyList(self, root: TreeNode) -> TreeNode:
        head, tail = self.genList(root)
        head.left, tail.right = tail, head
        return head



tree = TreeNode("[4,2,5,1,3]")
s = Solution()
head = s.treeToDoublyList(tree)

print(head.val, end=" ")
node = head.right
while node.val != head.val:
    print(node.val, end=" ")
    node = node.right

print()
print(head.val, end=" ")
node = head.left
while node.val != head.val:
    print(node.val, end=" ")
    node = node.left