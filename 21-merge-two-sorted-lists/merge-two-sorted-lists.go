/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if (list1 == nil && list2 == nil) {
        return nil
    } 
    if (list1 == nil){
        return list2
    }
    if (list2 == nil) {
        return list1
    }
    var head *ListNode = nil
    var result *ListNode = nil
    for(list1 != nil && list2 != nil) {
        if(head == nil) {
            head = new(ListNode)
            result = head
        } else {
            head.Next = new(ListNode)
            head = head.Next
        }
        if (list1.Val < list2.Val) {
            head.Val = list1.Val
            head.Next = nil
            list1 = list1.Next
        } else {
            head.Val = list2.Val
            head.Next = nil
            list2 = list2.Next
        }
    }
    for (list1 != nil) {
        head.Next = new(ListNode)
        head = head.Next
        head.Val = list1.Val
        head.Next = nil
        list1 = list1.Next
    }
    for (list2 != nil) {
        head.Next = new(ListNode)
        head = head.Next
        head.Val = list2.Val
        head.Next = nil
        list2 = list2.Next
    }
    return result
}