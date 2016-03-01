# coding: utf-8

class Solution:
    # @param digits, a list of integer digits
    # @return a list of integer digits
    def plusOne(self, digits):
        s = int(''.join([str(x) for x in digits]))
        return [int(x) for x in list(str(s+1))]

a = Solution()
print a.plusOne([0])