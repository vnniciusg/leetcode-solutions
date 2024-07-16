/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    numSet := make(map[int]struct{})
    for _, num := range nums {
        numSet[num] = struct{}{}
    }

    var prev *ListNode
    curr := head

    for curr != nil {
        if _, exists := numSet[curr.Val]; exists {
            if prev != nil {
                prev.Next = curr.Next
            } else {
                head = curr.Next
            }
        } else {
            prev = curr
        }
        curr = curr.Next
    }

    return head
}
