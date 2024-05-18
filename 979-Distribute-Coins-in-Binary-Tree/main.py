from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def distributeCoins(self, root: Optional[TreeNode]) -> int:
        def dfs(node: Optional[TreeNode]) -> tuple[int, int]:
            
            if node is None:
                return 0, 0
            
            left_excess, left_moves = dfs(node.left)
            right_excess, right_moves = dfs(node.right)
            
            excess = node.val + left_excess + right_excess - 1
            
            moves = left_moves + right_moves + abs(left_excess) + abs(right_excess)
            
            return excess, moves
        
        _, total_moves = dfs(root)
        return total_moves