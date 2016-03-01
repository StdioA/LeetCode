# coding: utf-8

class Solution:
	# @return a string
	def convert(self, s, nRows):
		if nRows == 1:
			return s

		chlist = []
		for x in range(2*nRows-1):
			chlist.append([])

		n = 0
		for x in s:
			chlist[n].append(x)
			n = (n+1)%(2*(nRows-1))

		ans = ''

		for row in range(nRows):
			if 0 < row < nRows-1:
				ano = 2*(nRows-1) - row
				pos = 1
				for x in chlist[ano]:
					chlist[row].insert(pos, x)
					pos += 2
			ans += ''.join(chlist[row])

		return ans

a = Solution()
print a.convert("PAYPALISHIRING", 3)

