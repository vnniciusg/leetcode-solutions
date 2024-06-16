/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

    if root == nil {
        return 0
    }

    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }

        leftNode := dfs(root.Left)
        rightNode := dfs(root.Right)

        return max(leftNode, rightNode) + 1
    }

    return dfs(root)
    
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
