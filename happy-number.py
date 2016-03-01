# coding: utf-8

class Solution:
    # @param {integer} n
    # @return {boolean}
    def isHappy(self, n):
        pre = {}
        while n != 1:
            if not pre.get(n, False):
                pre[n] = True
            else:
                return False
            l = [int(x) for x in str(n)]
            n = sum([x*x for x in l])
        return True

print Solution().isHappy(19)