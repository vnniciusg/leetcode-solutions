# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeInBetween(self, list1: ListNode, a: int, b: int, list2: ListNode) -> ListNode:
        
        curr_node = list1
        for _ in range(a - 1):
            curr_node = curr_node.next
        
        merge_end_node = curr_node.next
        for _ in range(b - a + 1):
            merge_end_node = merge_end_node.next
        
        curr_node.next = list2

        end_of_list2 = list2
        while end_of_list2.next:
            end_of_list2 = end_of_list2.next

        end_of_list2.next = merge_end_node

        return list1

        
