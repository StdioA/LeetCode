# coding: utf-8

# Definition for singly-linked list.
# class ListNode:
#	 def __init__(self, x):
#		 self.val = x
#		 self.next = None

from made_list import ListNode

class Solution:
	# @param two ListNodes
	# @return a ListNode
	def mergeTwoLists(self, l1, l2):
		pos1, pos2 = l1, l2

		if l1 == None:
			return l2
		elif l2 == None:
			return l1
		elif l1.val < l2.val:
			head = ListNode(l1.val)
			pos1 = pos1.next
		else:
			head = ListNode(l2.val)
			pos2 = pos2.next

		pre = head
		while pos1!=None and pos2!=None:
			if pos1.val < pos2.val:
				minNum = pos1.val
				pos1 = pos1.next
			else:
				minNum = pos2.val
				pos2 = pos2.next

			cur = ListNode(minNum)
			pre.next = cur
			pre = cur

		if pos1 == None:
			while pos2 != None:
				cur = ListNode(pos2.val)
				pre.next = cur
				pre = cur
				pos2 = pos2.next
		elif pos2 == None:
			while pos1 != None:
				cur = ListNode(pos1.val)
				pre.next = cur
				pre = cur
				pos1 = pos1.next
		return head

a = Solution()
l1 = ListNode('1->3->5->7')
l2 = ListNode('2->4->6->8')
print a.mergeTwoLists(l1, l2)
