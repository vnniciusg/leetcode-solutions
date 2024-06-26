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
    
    pub fn balance_bst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        let mut node_list = Vec::new();
        Self::inorder_traversal(&root, &mut node_list);
        Self::build_balanced_bst(&node_list, 0, node_list.len() as i32 - 1)
    }

    fn inorder_traversal(node: &Option<Rc<RefCell<TreeNode>>>, node_list: &mut Vec<i32>) {
        if let Some(node_ref) = node {
            let node = node_ref.borrow();
            Self::inorder_traversal(&node.left, node_list);
            node_list.push(node.val);
            Self::inorder_traversal(&node.right, node_list);
        }
    }

    fn build_balanced_bst(node_list: &[i32], start: i32, end: i32) -> Option<Rc<RefCell<TreeNode>>> {
        if start > end {
            return None;
        }
        let mid = start + (end - start) / 2;
        let node = Rc::new(RefCell::new(TreeNode::new(node_list[mid as usize])));
        node.borrow_mut().left = Self::build_balanced_bst(node_list, start, mid - 1);
        node.borrow_mut().right = Self::build_balanced_bst(node_list, mid + 1, end);
        Some(node)
    }
}
