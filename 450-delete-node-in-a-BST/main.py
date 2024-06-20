# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:

        def find_min(node: TreeNode) -> TreeNode:
            while node.left:
                node = node.left
            return node
        
        def dfs(root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
            if not root:
                return root

            if key > root.val:
                root.right = dfs(root.right, key)
            elif key < root.val:
                root.left = dfs(root.left, key)
            else:
                if not root.left:
                    return root.right
                if not root.right:
                    return root.left
                
                minNode = find_min(root.right)
                root.val = minNode.val
                root.right = dfs(root.right, minNode.val)
            
            return root
        
        return dfs(root, key)

