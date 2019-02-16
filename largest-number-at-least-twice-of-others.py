class Solution(object):
    def dominantIndex(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) < 2:
            return 0
        
        maxnum, another = nums[0], nums[1]
        index = 0
        if maxnum < another:
            maxnum, another = another, maxnum
            index = 1
        
        for tmpind, obj in enumerate(nums[2:]):
            tmpind += 2
            tmp = None
            if obj > maxnum:
                tmp, maxnum = maxnum, obj
                index = tmpind

                if tmp > another:
                    another = tmp
            else:
                another = max(another, obj)
        
        print(maxnum, another)
        if another * 2 <= maxnum:
            return index
        return -1


if __name__ == '__main__':
    s = Solution()
    print(s.dominantIndex([0,0,0,1]))
