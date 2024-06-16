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
    pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        
        if root.is_none() {
            return 0;
        }

        dfs(root)
    }
}


fn dfs(node: Option<Rc<RefCell<TreeNode>>>) -> i32 {
           
    if let Some(node) = node {
        let left_depth = dfs(node.borrow().left.clone());
        let right_depth = dfs(node.borrow().right.clone());
        return std::cmp::max(left_depth, right_depth) + 1;
    }
    
    0
}
