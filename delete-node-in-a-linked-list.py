# coding: utf-8

from made_list import ListNode

class Solution(object):
    def deleteNode(self, node):
        """
        :type node: ListNode
        :rtype: void Do not return anything, modify node in-place instead.
        """
        pre_node = node
        node = node.next
        while node.next:
            print(node.val)
            pre_node.val = node.val
            pre_node, node = pre_node.next, node.next

        pre_node.val = node.val
        pre_node.next = None

l = ListNode("1->2->3->4")
a = Solution()
a.deleteNode(l.next)
print(l)
