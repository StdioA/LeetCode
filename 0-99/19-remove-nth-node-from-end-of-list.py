# coding: utf-8

# Definition for singly-linked list.
# class ListNode:
#	 def __init__(self, x):
#		 self.val = x
#		 self.next = None

from made_list import ListNode

class Solution:
	# @return a ListNode
	def removeNthFromEnd(self, head, n):
		pre = pos = head
		if pos.next == None:
			return None

		while pos != None:
			temp = pos
			for x in range(n):
				temp = temp.next

			if temp == None:		#删除pos
				if pos == head:
					head = head.next
					del pos
				else:
					pre.next = pos.next
					del pos
				return head

			pre = pos
			pos = pos.next
		return head

a = Solution()

h = ListNode('1->2->3->4')

print a.removeNthFromEnd(h, 2)