# coding: utf-8

class Solution:
    # @param n, an integer
    # @return an integer
    def reverseBits(self, n):
        binstr = bin(n)[2:].zfill(32)
        ans = int(binstr[::-1], base=2)
        return ans

a = Solution()
print a.reverseBits(43261596)