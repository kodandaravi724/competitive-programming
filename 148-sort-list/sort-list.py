# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None:
            return None
        if head.next is None:
            return head
        slow, fast = head, head
        while fast is not None and fast.next is not None and fast.next.next is not None:
            slow = slow.next
            fast = fast.next.next
        temp = slow.next
        slow.next = None
        head1 = self.sortList(head)
        head2 = self.sortList(temp)
        sortedHead = ListNode()
        result = sortedHead
        prev = None
        while head1 is not None and head2 is not None:
            if head1.val < head2.val:
                sortedHead.val = head1.val
                sortedHead.next = ListNode()
                prev = sortedHead
                sortedHead = sortedHead.next
                head1 = head1.next
            else:
                sortedHead.val = head2.val
                sortedHead.next = ListNode()
                prev = sortedHead
                sortedHead = sortedHead.next
                head2 = head2.next
            
        while head1 is not None:
            sortedHead.val = head1.val
            head1 = head1.next
            sortedHead.next = ListNode()
            prev = sortedHead
            sortedHead = sortedHead.next
            
        while head2 is not None:
            sortedHead.val = head2.val
            head2 = head2.next
            sortedHead.next = ListNode()
            prev = sortedHead
            sortedHead = sortedHead.next
        prev.next = None
        return result
        
        