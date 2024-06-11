# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:

        if not lists:
            return None

        n = len(lists)
        
        ans = []
        for i in range(n):
            
            if not lists[i]:
                continue

            head = lists[i]
            while head is not None:
                ans.append(head.val)
                head = head.next

        if not ans:
            return None

        ans.sort()

        dummy = ListNode()
        current = dummy

        for val in ans:
            current.next = ListNode(val)
            current = current.next
        
        return dummy.next
        
