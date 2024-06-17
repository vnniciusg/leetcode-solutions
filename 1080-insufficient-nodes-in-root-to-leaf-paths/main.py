# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sufficientSubset(self, root: Optional[TreeNode], limit: int) -> Optional[TreeNode]:
        
        def dfs(node: Optional[TreeNode], sum_so_far: int) -> Optional[TreeNode]:
            if not node:
                return None
            
            sum_so_far += node.val
            
            if not node.left and not node.right:
                if sum_so_far < limit:
                    return None
                else:
                    return node
            
            node.left = dfs(node.left, sum_so_far)
            node.right = dfs(node.right, sum_so_far)
            
            if not node.left and not node.right:
                return None
            
            return node

        return dfs(root, 0)



        