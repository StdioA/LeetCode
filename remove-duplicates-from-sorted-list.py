# Definition for singly-linked list.
# class ListNode:
#	 def __init__(self, x):
#		 self.val = x
#		 self.next = None

from made_list import ListNode

class Solution:
	# @param head, a ListNode
	# @return a ListNode
	def deleteDuplicates(self, head):
		if head == None:
			return head
		cur = head
		while cur.next != None:
			if cur.val == cur.next.val:
				cur.next = cur.next.next
			else:
				cur = cur.next
		return head

a = Solution()
l = ListNode('1->1->2')
print a.deleteDuplicates(l)