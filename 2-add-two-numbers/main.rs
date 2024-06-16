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
    pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        
        let mut head = Some(Box::new(ListNode::new(0)));
        let mut curr = head.as_mut();
        let mut carry = 0;

        let (mut l1, mut l2) = (l1, l2);

        while l1.is_some() || l2.is_some() || carry != 0 {
            let v1 = if let Some(ref node) = l1 { node.val } else { 0 };
            let v2 = if let Some(ref node) = l2 { node.val } else  { 0 };

            let sum  = v1 + v2 + carry;
            let digit = sum % 10;
            carry = sum / 10;

            if let Some(node) = curr {
                node.next = Some(Box::new(ListNode::new(digit)));
                curr = node.next.as_mut();
            }
 
            l1 = if let Some(node) = l1 { node.next } else { None };
            l2 = if let Some(node) = l2 { node.next } else { None };
        } 

        head.unwrap().next
    }
}
