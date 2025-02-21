/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    seen map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
    fe := FindElements{seen: make(map[int]struct{})}
    var dfs func(node *TreeNode, val int)
    dfs = func(node *TreeNode, val int) {
        if node == nil {
            return
        }

        fe.seen[val] = struct{}{}
        dfs(node.Left, val * 2 + 1)
        dfs(node.Right, val * 2 + 2)
    }

    dfs(root, 0)
    return fe
}


func (this *FindElements) Find(target int) bool {
    _, exists := this.seen[target]
    return exists
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
