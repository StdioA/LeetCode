class Solution:
    def simplifyPath(self, path: str) -> str:
        segments = path.split("/")
        stack = []
        for segment in segments[1:]:
            if segment in ("", "."):
                continue
            elif segment == "..":
                if len(stack) > 0:
                    stack.pop(-1)
            else:
                stack.append(segment)
        return "/" + "/".join(stack)
