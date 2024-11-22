/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getNext(head *ListNode) *ListNode  {
    for head != nil && head.Next != nil && head.Val == head.Next.Val {
    {
        for head != nil && head.Next != nil && head.Val == head.Next.Val {
        head = head.Next
        }
    head = head.Next
    }
    }
    return head
}

func deleteDuplicates(head *ListNode) *ListNode {
    var nodeHead, currentPointer *ListNode = nil, nil
    for head != nil {
        head = getNext(head)
        if nodeHead == nil {
            nodeHead = head
            currentPointer = head
        } else {
            currentPointer.Next = head
            currentPointer = head
        }
        if head != nil {
            head = head.Next
        }
    }
    return nodeHead
}