# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapNodes(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        
        def get_length(node: ListNode) -> int:
            length = 0
            while node:
                length += 1
                node = node.next
            return length
        

        length = get_length(head)

        if k > length:
            return None
        

        kth_node_from_beginning = head
        for _ in range(k-1):
            kth_node_from_beginning = kth_node_from_beginning.next
        

        kth_node_from_end = head
        for _ in range(length - k):
            kth_node_from_end = kth_node_from_end.next
        
        kth_node_from_beginning.val, kth_node_from_end.val = kth_node_from_end.val, kth_node_from_beginning.val

        return head

