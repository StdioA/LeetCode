# coding: utf-8

#    1
#   / \
#  2   3
#     /
#    4
#     \
#      5
# The above binary tree is serialized as "{1,2,3,#,#,4,#,#,5}"


class TreeNode(object):
    def __init__(self, x=None):
        if isinstance(x, str):
            nums = x[1:-1].split(',')                       # split the params
            self.val = int(nums[0].strip())
            self.left = None
            self.right = None

            nodelist = [self]                               # The nodelist records the nodes that is going to be executed
            pointer = self
            is_left = True

            for num in nums[1:]:
                num = num.strip()
                if num != '#':
                    temp = TreeNode(int(num))
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

        else:
            self.val = x
            self.left = None
            self.right = None

    def __str__(self):
        if not self:
            return "{}"

        pre = []
        cur = [self]
        tra = []
        while cur:
            tra.append([])
            pre, cur = cur[:], []
            for node in pre:
                if not isinstance(node, str):
                    tra[-1].append(node.val)
                    if node.left:
                        cur.append(node.left)
                    else:
                        cur.append("#")
                    if node.right:
                        cur.append(node.right)
                    else:
                        cur.append("#")
                else:
                    tra[-1].append("#")

        s = "{"
        for line in tra[:-1]:
            for d in line:
                s += str(d)+", "
        s = s[:-2]+"}"
        return s

    def __repr__(self):
        return self.__str__()


if __name__ == '__main__':
    a = TreeNode("{1, 2, 3, #, #, 4, #, #, 5}")
    print(a)
