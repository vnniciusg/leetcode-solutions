# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def balanceBST(self, root: TreeNode) -> TreeNode:
        
        node_list: list[int] = []

        def traverse(node: TreeNode):
            if not node: return

            traverse(node.left)
            node_list.append(node.val)
            traverse(node.right)

        def build_bst(left: int, right: int) -> TreeNode:
            if left > right: return None
            mid: int = left + (right - left) // 2
            root = TreeNode(node_list[mid])
            root.left = build_bst(left, mid-1)
            root.right = build_bst(mid + 1, right)
            return root
        
        traverse(root)
        return build_bst(0, len(node_list) - 1)
            
