# coding: utf-8

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

from made_list import ListNode

class Solution:
    # @param {ListNode} head
    # @param {integer} val
    # @return {ListNode}
    def removeElements(self, head, val):
        if not head:
            return head

        while head and head.val == val:
            temp = head
            head = head.next
            del temp

        cur = head

        if cur:
            while cur.next:
                if cur.next.val == val:
                    temp = cur.next
                    cur.next = cur.next.next
                    del temp
                else:
                    cur = cur.next

        return head