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
    pub fn merge_in_between(list1: Option<Box<ListNode>>, a: i32, b: i32, list2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        
        let mut dummy = Some(Box::new(ListNode { val: 0, next: list1 }));

        let mut curr_node = &mut dummy;
        for _ in 0..a {
            curr_node = &mut curr_node.as_mut().unwrap().next;
        }

        let mut merge_end_node = curr_node.clone();
        for _ in 0..(b - a + 2) {
            merge_end_node = merge_end_node.unwrap().next.take();
        }

        curr_node.as_mut().unwrap().next = list2;
        while curr_node.as_mut().unwrap().next.is_some() {
            curr_node = &mut curr_node.as_mut().unwrap().next;
        }
        
        curr_node.as_mut().unwrap().next = merge_end_node;
        dummy.unwrap().next
    }
}