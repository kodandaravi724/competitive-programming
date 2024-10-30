# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if head is None or head.next is None:
            return False
        slowPointer, fastPointer = head, head
        slowPointer = slowPointer.next
        fastPointer = fastPointer.next.next
        while fastPointer is not None and fastPointer.next is not None and slowPointer != fastPointer:
            slowPointer = slowPointer.next
            fastPointer = fastPointer.next.next
        if fastPointer is not None and fastPointer.next is not None:
            return True
        return False 
        
        