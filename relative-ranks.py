class Solution(object):
    def findRelativeRanks(self, nums):
        """
        :type nums: List[int]
        :rtype: List[str]
        """
        sorted_nums = list(sorted(nums, reverse=True))

        def transform_rank(rank):
            ranks = {
                0: "Gold Medal",
                1: "Silver Medal",
                2: "Bronze Medal"
            }
            return ranks.get(rank, str(rank + 1))

        return [transform_rank(sorted_nums.index(i)) for i in nums]


s = Solution()
print(s.findRelativeRanks([5, 4, 3, 2, 1]))
