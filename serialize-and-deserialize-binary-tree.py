# coding: utf-8

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

from made_bitree import TreeNode

class Codec:

    def serialize(self, root):                                                  # 由binary-tree-level-order-traversal修改
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        if not root:
            return "{}"

        pre = []
        cur = [root]
        tra = []
        if root:
            while cur:
                tra.append([])
                pre, cur = cur[:], []
                for node in pre:
                    if not isinstance(node, str):
                        tra[-1].append(node.val)
                        if node.left:
                            cur.append(node.left)
                        else:
                            cur.append("null")
                        if node.right:
                            cur.append(node.right)
                        else:
                            cur.append("null")
                    else:
                        tra[-1].append("null")

        s = "{"
        for line in tra[:-1]:
            for d in line:
                s += str(d)+","
        s = s[:-1]+"}"
        return s

    def deserialize(self, data):                                                # 由made-bitree修改
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        if not data or data == "{}":
            return None
        
        nums = data[1:-1].split(',')                    # split the params
        root = TreeNode(int(nums[0].strip()))
        root.left = None
        root.right = None
        
        nodelist = [root]                               # The nodelist records the nodes that is going to be executed
        pointer = root
        is_left = True

        for num in nums[1:]:
            if num != 'null':
                temp = TreeNode(int(num.strip()))
                nodelist.append(temp)
                if is_left:                             # put the node in the proper place (the L/R of the current node)
                    pointer.left = temp
                else:
                    pointer.right = temp
                                                        # the character '#' means "pass"
            if not is_left:                             # If the value is_left is True, it means that the current node
                nodelist = nodelist[1:]                 # had been executed (the two sons are both valued or skipped
                pointer = nodelist[0]                   # by the operation), so the pointer should point the next node
            is_left = not is_left                       # which is going to be executed.

        return root

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
codec = Codec()
tree = TreeNode("{1,2,3,#,#,4,#,#,5}")
print(codec.serialize(tree))
print(codec.serialize(codec.deserialize(codec.serialize(tree))))
