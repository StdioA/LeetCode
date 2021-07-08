from typing import List

class Solution:
    def lex(self, s):
        p = s[0]
        is_lex, strict = True, True
        for c in s[1:]:
            if c < p:
                is_lex, strict = False, False
                break
            elif c == p:
                strict = False
            p = c
        return is_lex, strict

    def minDeletionSize(self, strs: List[str]) -> int:
        print(strs)
        transposed_strs = ["".join(col) for col in zip(*strs)]
        count = 0
        print(transposed_strs)
        for col in transposed_strs:
            is_lex, strict = self.lex(col)
            print(col, is_lex, strict)
            if not is_lex:
                count += 1
            elif strict:
                break
        return count


a = Solution()
print(a.minDeletionSize(["xga","xfb","yfa"]))