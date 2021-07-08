# coding: utf-8

class Solution:
    # @param s, a string
    # @return a boolean
    def isPalindrome(self, s):
    	k = ''
    	for c in s:
    		if 'a'<=c<='z' or 'A'<=c<='Z' or '0'<=c<='9':
    			k += c
    	k = k.lower()
    	return k[::-1]==k

a = Solution()
print a.isPalindrome('A man, a plan, a canl: Panama')