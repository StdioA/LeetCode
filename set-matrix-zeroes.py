# coding: utf-8

class Solution:
    # @param {integer[][]} matrix
    # @return {void} Do not return anything, modify matrix in-place instead.
    def setZeroes(self, matrix):
        cols = set()
        rows = set()

        for i, row in enumerate(matrix):
            for j, num in enumerate(row):
                if num == 0:
                    cols.add(i)
                    rows.add(j)

        for i in cols:
            for j in range(len(matrix[0])):
                matrix[i][j] = 0
        for j in rows:
            for i in range(len(matrix)):
                matrix[i][j] = 0

a = Solution()
l = [[0,1]]
a.setZeroes(l)
print l
