/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    
    length := getLength(head)

    if k > length {
        return &ListNode{}
    }

    kThBeggin := head
    for range(k-1) {
        kThBeggin = kThBeggin.Next
    }

    kThEnd := head
    for range(length - k) {
        kThEnd = kThEnd.Next
    }

    kThBeggin.Val, kThEnd.Val = kThEnd.Val, kThBeggin.Val
    return head
}

func getLength(node *ListNode) int {
        var length int = 0
        for node != nil {
            length ++
            node = node.Next
        }

        return length
    }

