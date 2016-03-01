# coding: utf-8

class Solution:
    # @param {string} s
    # @param {string} t
    # @return {boolean}
    def isIsomorphic(self, s, t):
        d = {}
        for i in range(len(s)):
            c1, c2 = s[i], t[i]
            k = d.get(c1,None)
            if k and c2!=k:
                return False
            elif not k:
                if c2 in d.values():
                    return False
            d[c1] = c2
        return True

print Solution().isIsomorphic('efg','add')