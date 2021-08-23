class Solution(object):
    def pathSum(self, root, sum_):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: List[List[int]]
        """
        
        solutions = []
        
        def find(node_list):
            node = node_list[-1]
            if not any([node.left, node.right]):
                path_sum = sum(n.val for n in node_list)
                if path_sum == sum_:
                    solutions.append([n.val for n in node_list])
            
            if node.left:
                find(node_list + [node.left])
            
            if node.right:
                find(node_list + [node.right])
            return
        
        if not root:
            return []

        find([root])
        return solutions
