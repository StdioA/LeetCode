# coding: utf-8

class Solution(object):
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        nums = [int(x) for x in str(num)][::-1]
        chars = [["I","V"],["X","L"],["C","D"],["M"]]
        ansl = []
        for i, num in enumerate(nums):
            if num <= 3:
                cnum = chars[i][0]*num
            elif num == 4:
                cnum = chars[i][0] + chars[i][1]
            elif 5 <= num <= 8:
                cnum = chars[i][1] + chars[i][0]*(num-5)
            else:
                cnum = chars[i][0] + chars[i+1][0]
            ansl.append(cnum)
            
        return "".join(reversed(ansl))
        
