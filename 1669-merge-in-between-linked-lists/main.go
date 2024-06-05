package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

    currNode := list1
    for range(a-1){
        currNode = currNode.Next
    }

    mergeEndNode := currNode.Next
    for range( b - a + 1){
        mergeEndNode = mergeEndNode.Next
    }

    currNode.Next = list2

    endList2 := list2
    for endList2.Next != nil {
        endList2 = endList2.Next
    }

    endList2.Next = mergeEndNode
    return list1
}
