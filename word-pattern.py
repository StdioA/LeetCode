# coding: utf-8

class Solution(object):
    def wordPattern(self, pattern, str):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """
        strl = str.split(" ")
        words = {}
        
        if len(pattern) != len(strl):
            return False
        
        for i, p in enumerate(pattern):
            if p not in words:
                words[p] = strl[i]
            elif words[p] != strl[i]:
                return False
                
            if len(filter(lambda w: w[1]==strl[i], words.items())) != 1:
                return False
        
        return True
        
