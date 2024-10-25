from heapq import *

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        minHeap = []
        for i in range(len(lists)):
            if lists[i] is not None:
                heappush(minHeap, (lists[i].val, i))
                lists[i] = lists[i].next
        if minHeap:
            minValue, index = heappop(minHeap)
            head = ListNode(minValue)
            if lists[index] is not None:
                heappush(minHeap, (lists[index].val, index))
                lists[index] = lists[index].next
            cur = head
            while minHeap:
                minValue, index = heappop(minHeap)
                cur.next = ListNode(minValue)
                cur = cur.next
                if lists[index] is not None:
                    heappush(minHeap, (lists[index].val, index))
                    lists[index] = lists[index].next
            return head
        else:
            return None


        