# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def getDirections(self, root: Optional[TreeNode], startValue: int, destValue: int) -> str:
        
        def dfs(node: Optional[TreeNode], val: int, path: List[str]) -> bool:

            if node.val == val:
                return True
            
            if node.left and dfs(node.left, val, path):
                path += "L"
            elif node.right and dfs(node.right, val, path):
                path += "R"
            
            return path

        s, d = [], []
        dfs(root, startValue, s)
        dfs(root, destValue, d)

        while len(s) and len(d) and s[-1] == d[-1]:
            s.pop()
            d.pop()
        
        return "".join("U" * len(s)) + "".join(reversed(d))

