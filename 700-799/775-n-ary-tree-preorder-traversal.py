"""
# Definition for a Node.
class Node:
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution:
    def preorder(self, root: 'Node') -> List[int]:
        if root is None:
            return []
        stack = [root]
        result = []
        while stack:
            node = stack.pop(0)
            result.append(node.val)
            stack = node.children + stack
        return result
