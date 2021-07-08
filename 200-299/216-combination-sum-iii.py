# coding: utf-8

class Solution(object):
    def combinationSum3(self, k, n):
        """
        :type k: int
        :type n: int
        :rtype: List[List[int]]
        """
        import itertools
        it = itertools.combinations(range(1, 10), k)
        filted = filter(lambda x: sum(x) == n, it)
        ans = [list(x) for x in filted]

        return ans

a = Solution()
print a.combinationSum3(3, 9)
