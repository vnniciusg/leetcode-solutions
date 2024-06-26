/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    var node_list []int
    traverse(root, &node_list)
    return buildBST(0, len(node_list)-1, node_list)
}

func traverse(node *TreeNode, node_list *[]int) {
    if node == nil {
        return
    }
    traverse(node.Left, node_list)
    *node_list = append(*node_list, node.Val)
    traverse(node.Right, node_list)
}

func buildBST(left, right int, node_list []int) *TreeNode {
    if left > right {
        return nil
    }
    mid := left + (right-left)/2
    root := &TreeNode{Val: node_list[mid]}
    root.Left = buildBST(left, mid-1, node_list)
    root.Right = buildBST(mid+1, right, node_list)
    return root
}
