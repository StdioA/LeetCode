# coding: utf-8
# UNFINISHED!

class Solution(object):
    def maximalRectangle(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if (not matrix) or (not matrix[0]):
            return 0

        height, length = len(matrix), len(matrix[0])
        rec = [[[0, 0] for _ in range(length)] for _ in range(height)]

        for i,line in enumerate(matrix):
            for j, c, in enumerate(line):
                if c == "1":
                    if i == 0:
                        if j > 0 and matrix[i][j-1] == "1":
                            rec[i][j][1] = rec[i][j-1][1] + 1
                            rec[i][j][0] = 1
                        else:
                            rec[i][j] = [1, 1]
                    elif j == 0:
                        if i > 0 and matrix[i-1][j] == "1":
                            rec[i][j][0] = rec[i-1][j][01] + 1
                            rec[i][j][1] = 1
                        else:
                            rec[i][j] = [1, 1]
                    else:
                        if 

        print rec


a = Solution()
print a.maximalRectangle([["0", "1", "0"], ["0", "1", "0"], ["0", "0", "1"]])
# print a.maximalRectangle(["0"])
