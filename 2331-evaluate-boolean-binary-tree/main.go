package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}

	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)

	switch root.Val {
	case 2:
		return left || right
	case 3:
		return left && right
	default:
		return false
	}
}
