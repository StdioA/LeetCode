# coding: utf-8

class Solution(object):
    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        if not digits:
            return []
            
        import itertools
        
        ss = [None, None, "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"]
        
        strs = [ss[int(d)] for d in digits]
        it = itertools.product(*strs)
        
        return ["".join(o) for o in it]
        
a = Solution()
print a.letterCombinations("23")
