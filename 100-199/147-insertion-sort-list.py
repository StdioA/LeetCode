# coding: utf-8

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None
from made_list import ListNode

class Solution(object):
    def insertionSortList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head:
            return None

        if not head.next:
            return head

        pointer = head.next
        lhead = ListNode(head.val)

        while pointer:
            val = pointer.val
            node = ListNode(val)

            if val < lhead.val:        # 插入现有头部之前
                node.next = lhead
                lhead = node
            else:
                p = lhead
                while p.next and (val > p.next.val):
                    p = p.next

                p.next, node.next = node, p.next

            pointer = pointer.next

        return lhead

a = Solution()
# head = ListNode("4->1->5->2->3")
head = ListNode(4)
print a.insertionSortList(head)
        
