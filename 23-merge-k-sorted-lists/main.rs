// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
// 
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
        
        if lists.is_empty() {
            return None;
        }

        let mut ans: Vec<i32> = Vec::new();
        for list in lists {
            let mut head = list;
            while let Some(node) = head {
                ans.push(node.val);
                head = node.next;
            }
        }

        if ans.is_empty() {
            return None;
        }

        ans.sort();

        let mut dummy = ListNode::new(0);
        let mut current = &mut dummy;

        for val in ans {
            current.next = Some(Box::new(ListNode::new(val)));
            current = current.next.as_mut().unwrap();
        }

        dummy.next
    }
}
