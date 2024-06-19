/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    
    if p == nil && q == nil {
        return true
    }

    if (p == nil && q != nil) || (q == nil && p != nil) {
        return false
    }

    if p.Val != q.Val {
        return false
    }

    return isSameTree(q.Right, p.Right) && isSameTree(q.Left, p.Left)
}
