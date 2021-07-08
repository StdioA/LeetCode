# coding: utf-8

class Solution(object):
    def minimumTotal(self, triangle):
        """
        :type triangle: List[List[int]]
        :rtype: int
        """
        length = len(triangle)
        for i in reversed(range(length-1)):
            for j in range(i+1):
                triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])

        return triangle[0][0]

a = Solution()
print a.minimumTotal([[2],[3,4],[6,5,7],[4,1,8,3]])
