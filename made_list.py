# coding: utf-8

class ListNode(object):
	def __init__(self, x):
		if isinstance(x, str):
			lEle = x.split('->')
			self.val = int(lEle[0])

			pre = self
			for ele in lEle[1:]:
				temp = ListNode(int(ele))
				pre.next = temp
				pre = pre.next
		else:
			self.val = x
			self.next = None

	def __str__(self):
		pos = self
		s = '{'
		while pos != None:
			s += str(pos.val)+','
			pos = pos.next

		s = s[:-1]
		s += '}'
		return s

	def __repr__(self):
		pos = self
		s = ''
		while pos != None:
			s += str(pos.val)+'->'
			pos = pos.next

		s = s[:-2]
		return s