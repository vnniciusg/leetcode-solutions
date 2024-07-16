# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def modifiedList(self, nums: List[int], head: Optional[ListNode]) -> Optional[ListNode]:
        
        if not head: return head

        curr = head
        prev = None

        nums = set(nums)

        while curr:
            if curr.val in nums:
                if prev:
                    prev.next = curr.next
                else:
                    head = curr.next
            else:
                prev = curr

            curr = curr.next

        return head
