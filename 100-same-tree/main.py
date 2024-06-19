from typing import Optional

class TreeNode:
	def __init__(self, val = 0, right, left):
		self.val = val
		self.right = right
		self.left = left


class Solution:
	def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:

	if not q and not p:
		return True

	if (not q and p) or (not p and q):
		return False
	
	if p.val != q.val:
		return False

	return self.isSameTree(q.right, p.right) and self.isSameTree(q.left, r.left)
