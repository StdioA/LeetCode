from collections import defaultdict

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        maxLength = 0
        freqMap = defaultdict(int)
        i, j = 0, 0
        
        while j < len(s):
            # move j forward
            while j < len(s) and freqMap[s[j]] < 1:
                freqMap[s[j]] += 1
                j += 1
                
            maxLength = max(maxLength, j - i)
            # move i forward
            print(i, j, freqMap)
            while j < len(s) and i < j and freqMap[s[j]] >= 1:
                freqMap[s[i]] -= 1
                i += 1
      
        return maxLength


s = Solution()
print(s.lengthOfLongestSubstring("abcabcbb"))
