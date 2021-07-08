# coding: utf-8

class Solution:
	# @param    A       a list of integers
	# @param    elem    an integer, value need to be removed
	# @return an integer
	def removeElement(self, A, elem):
		if A == []:
			return 0
		m = 0
		for index in range(len(A)):
			if A[index] == elem:
				m += 1
			else:
				A[index-m] = A[index]

		print A[:index-m+1]
		return index-m+1

a = Solution()
print a.removeElement([1,3,1,3,3], 3)