# coding: utf-8
# UNFINISHED

class Solution(object):
    def groupAnagrams_TLE(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """
        # 这个算法也是错的！
        anas = []

        for s in strs:
            s_set = set(s)
            for i, t in enumerate(anas):
                c_set = t[0]

                if c_set == s_set:
                    anas[i][1].append(s)
                    break
            else:
                anas.append([s_set, [s]])

        anas = [o[1] for o in anas]
        map(lambda x: x.sort(), anas)

        return anas


a = Solution()
print a.groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
