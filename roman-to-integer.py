# coding: utf-8

class Solution:
    # @return an integer
    def romanToInt(self, s):
        chnum = {
        		'I' : 1,
        		'V' : 5,
        		'X' : 10,
        		'L' : 50,
        		'C' : 100,
        		'D' : 500,
        		'M' : 1000,
        		}
        count = 0
        cur = s[0]
        num = 0
        chlist = []
        #统计每个字母及其出现的次数
        for x in s:
        	if x == cur:
        		count += 1
        	else:
        		chlist.append((cur, count))
        		count = 1
        		cur = x
        else:
        	chlist.append((cur, count))
        #好像是第一次用for-else
        #根据字母顺序累加数
        for index in range(len(chlist)-1):
        	ch, count = chlist[index]
        	if chnum[chlist[index][0]] < chnum[chlist[index+1][0]]:		#如果前面的数比后面的数小(IIV, IX)，就说明应该用减
        		num -= chnum[ch]*count
        	else:														#前面比后面大用加(VII, LXX)
        		num += chnum[ch]*count
        #若前后两个字母不在同一数量级(IX, XI), 同样符合以上规则
        num += chnum[chlist[-1][0]] * chlist[-1][1]						#最后一个数需要加一下（没有什么需要跟它比，直接加上）
        return num

a = Solution()
print a.romanToInt('MDCCLXXVI')