# coding: utf-8

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

from made_list import ListNode

class Solution:
    # @param head, a ListNode
    # @return a boolean
    def hasCycle(self, head):
        nodes = {}
        cur = head
        while cur:
            if nodes.get(id(cur),None):
                return True
            else:
                nodes[id(cur)] = True
            cur = cur.next
        return False