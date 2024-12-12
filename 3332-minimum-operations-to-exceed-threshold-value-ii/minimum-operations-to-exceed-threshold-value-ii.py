from heapq import *

class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        minHeap = []
        for num in nums:
            heappush(minHeap, num)
        numOperations = 0
        while len(minHeap) >= 2:
            if minHeap[0] >= k:
                return numOperations
            minV = heappop(minHeap)
            maxV = heappop(minHeap)
            heappush(minHeap, 2*minV + maxV)
            numOperations+=1
        if minHeap[0] >= k:
            return numOperations
        return 0
