class Solution:
    def frequencySort(self, s: str) -> str:
        from collections import Counter
        c = Counter(s)
        return "".join([c * n for c, n in c.most_common()])


print(Solution().frequencySort('tree'))
