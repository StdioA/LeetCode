# coding: utf-8

from typing import List
from functools import reduce

#找一坨字符串的最长公共前缀
#减少比较次数
class Solution:
    # @return a string
    def longestCommonPrefix(self, strs):
        if len(strs) == 0:
            return ''
        longest = reduce(min, map(len, strs))
        if longest == 0:
            return ''

        for i in range(len(strs)-1):
            for x in range(longest):
                if strs[i][x] != strs[i+1][x]:
                    longest = x
                    break

        return strs[0][:longest]

    def longestCommonPrefixNew(self, strs: List[str]) -> str:
        prefix = strs[0]
        for s in strs[1:]:
            length = min(len(prefix), len(s))
            prefix = prefix[:length]
            for i in range(length-1, -1, -1):
                if prefix[i] != s[i]:
                    prefix = prefix[:i]
        return prefix


a = Solution()
s = ['a', 'ab']
print(a.longestCommonPrefixNew(s))
