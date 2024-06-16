# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        
        if not root:
            return 0

        def dfs(node: Optional[TreeNode]) -> int:
            
            if not node:
                return 0

            left_node = dfs(node.left)
            right_node = dfs(node.right)

            return max(left_node, right_node) + 1
        
        return dfs(root)

        
