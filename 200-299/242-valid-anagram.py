# coding: utf-8

class Solution(object):
    def isAnagram(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        d = {}
        for c in s:
            if c in d:
                d[c] += 1
            else:
                d[c] = 1

        for c in t:
            if not (c in d):
                return False
            elif d[c] == 0:
                return False
            else:
                d[c] -= 1

        for i in d.values():
            if i != 0:
                return False

        return True

a = Solution()
print a.isAnagram("anagram", "nagaram")
print a.isAnagram("ab", "b")
