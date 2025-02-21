# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class FindElements:

    def __init__(self, root: Optional[TreeNode]):
        self.seen = set()
        self._dfs(root, 0)

    def find(self, target: int) -> bool:
        return target in self.seen


    def _dfs(self, curr_node: Optional[TreeNode], curr_val: int) -> None:

        if curr_node is None: return

        self.seen.add(curr_val)
        self._dfs(curr_node.left, curr_val * 2 + 1)
        self._dfs(curr_node.right, curr_val * 2 + 2)
        


# Your FindElements object will be instantiated and called as such:
# obj = FindElements(root)
# param_1 = obj.find(target)
