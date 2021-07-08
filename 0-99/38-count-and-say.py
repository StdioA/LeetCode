# coding: utf-8

class Solution:
    # @return a string
    def countAndSay(self,n):
        cs = ['1']
        for i in range(1,n):
            count = 0
            s = cs[-1]
            num = s[0]
            snext = ''
            l = len(s)
            x = 0
            while x < l:
                if s[x] == num:
                    count += 1
                    x += 1
                else:
                    snext += str(count)+str(num)
                    num = s[x]
                    count = 1
                    x += 1
    
            snext += str(count)+str(num)
            cs.append(snext)
            
        return cs[-1]