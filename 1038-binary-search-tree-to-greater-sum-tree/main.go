/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
    
    reverseInorderTraversal(root, 0)
    return root
}


func reverseInorderTraversal(node *TreeNode, sum int) int {
    if node != nil {
        sum = reverseInorderTraversal(node.Right, sum)
        sum += node.Val
        node.Val = sum
        sum = reverseInorderTraversal(node.Left, sum)
    }
    return sum
}
