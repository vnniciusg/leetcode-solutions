package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	return dfs(root, target)
}

func dfs(node *TreeNode, target int) *TreeNode {

	if node == nil {
		return nil
	}

	node.Left = dfs(node.Left, target)
	node.Right = dfs(node.Right, target)

	if node.Left == nil && node.Right == nil && node.Val == target {
		return nil
	}

	return node
}
