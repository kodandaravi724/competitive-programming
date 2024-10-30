from heapq import *
class Solution:
    def findMaximizedCapital(self, k: int, w: int, profits: List[int], capital: List[int]) -> int:
        minHeap = []
        maxHeap = []
        finalCapital = w
        for i in range(len(capital)):
            heappush(minHeap, (capital[i], profits[i]))
        while k > 0:
            while minHeap and w >= minHeap[0][0]:
                capital, profit = heappop(minHeap)
                heappush(maxHeap, (-profit, capital))
            if maxHeap:
                # print(maxHeap)
                finalCapital += (-maxHeap[0][0])
                heappop(maxHeap)
                w = finalCapital
            k-=1
        return finalCapital

        