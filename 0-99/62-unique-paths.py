class Solution(object):
    def uniquePaths(self, m, n):
        """
        :type m: int
        :type n: int
        :rtype: int
        """
        # C(m+n-2, m - 1)
        m, n = m + n - 2, m - 1
        result = 1
        for i in range(m - n + 1, m + 1):
            result *= i
        for i in range(1, n + 1):
            result /= i
        return result
