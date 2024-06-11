/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    
    if lists == nil {
        return nil
    }

    ans := []int{}
    for _, list := range lists {
        head := list
        for head != nil {
            ans = append(ans, head.Val)
            head = head.Next
        }
    }

    if len(ans) == 0 {
        return nil
    }

    sort.Slice(ans, func(i, j int) bool {
        return ans[i] < ans[j]
    })

    dummy := &ListNode{Val: 0, Next: nil}
    current := dummy

    for _, val := range ans {
        current.Next = &ListNode{Val: val}
        current = current.Next
    }

    return dummy.Next
}
