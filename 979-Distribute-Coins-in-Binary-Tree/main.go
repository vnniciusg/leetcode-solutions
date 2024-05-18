package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	_, moves := dfs(root)
	return moves
}

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	left_excess, left_moves := dfs(node.Left)
	right_excess, right_moves := dfs(node.Right)

	excess := node.Val + left_excess + right_excess - 1
	moves := left_moves + right_moves + abs(left_excess) + abs(right_excess)

	return excess, moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
