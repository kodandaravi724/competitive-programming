from collections import defaultdict
from heapq import *

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = defaultdict(int)
        for n in nums:
            freq[n] += 1
        maxHeap = []
        for k1, v in freq.items():
            if len(maxHeap) < k or maxHeap[0][0] < v:
                heappush(maxHeap, (v, k1))
            if len(maxHeap) > k:
                heappop(maxHeap)
        elements = []
        while maxHeap:
            elements.append(maxHeap.pop()[1])
        return elements
        