// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
// 
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn bst_to_gst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        let mut sum = 0;
        Solution::reverse_inorder_traversal(&root, &mut sum);
        root
    }

    fn reverse_inorder_traversal(node: &Option<Rc<RefCell<TreeNode>>>, sum: &mut i32) {
        if let Some(n) = node {
            let mut node_borrow = n.borrow_mut();
            Solution::reverse_inorder_traversal(&node_borrow.right.clone(), sum);
            *sum += node_borrow.val;
            node_borrow.val = *sum;
            Solution::reverse_inorder_traversal(&node_borrow.left.clone(), sum);
        }
    }
}
