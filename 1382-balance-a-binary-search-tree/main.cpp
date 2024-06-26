/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* balanceBST(TreeNode* root) {
        vector<int> node_list;
        inorderTraversal(root, node_list);
        return buildBalancedBST(node_list, 0, node_list.size() - 1);
    }

private:
    void inorderTraversal(TreeNode* root, vector<int>& node_list) {
        if (!root) return;
        inorderTraversal(root->left, node_list);
        node_list.push_back(root->val);
        inorderTraversal(root->right, node_list);
    }

    TreeNode* buildBalancedBST(const vector<int>& node_list, int start, int end) {
        if (start > end) return nullptr;
        int mid = start + (end - start) / 2;
        TreeNode* node = new TreeNode(node_list[mid]);
        node->left = buildBalancedBST(node_list, start, mid - 1);
        node->right = buildBalancedBST(node_list, mid + 1, end);
        return node;
    }
};
