from operator import itemgetter

class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=itemgetter(0))
        interval_dict = {}

        for start, end in intervals:
            # Find closet
            closest_start = None
            for k in interval_dict:
                if k <= start:
                    if closest_start is None or closest_start < k:
                        closest_start = k
            if closest_start is None:
                closest_start = start
                interval_dict.setdefault(closest_start, end)

            istart, iend = closest_start, interval_dict[closest_start]
            if iend >= start:
                interval_dict[closest_start] = max(iend, end)
            else:
                interval_dict[start] = end
        
        return [list(i) for i in interval_dict.items()]
