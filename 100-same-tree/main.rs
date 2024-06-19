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
    pub fn is_same_tree(p: Option<Rc<RefCell<TreeNode>>>, q: Option<Rc<RefCell<TreeNode>>>) -> bool {
        match (p, q) {
            (Some(node_p), Some(node_q)) =>{
                let p_node_borrow = node_p.borrow();
                let q_node_borrow = node_q.borrow();
                p_node_borrow.val == q_node_borrow.val 
                    && Self::is_same_tree(p_node_borrow.left.clone(), q_node_borrow.left.clone())
                    && Self::is_same_tree(p_node_borrow.right.clone(), q_node_borrow.right.clone())
            }
            (None, None) => true,
            _ => false,
        }

    }
}
