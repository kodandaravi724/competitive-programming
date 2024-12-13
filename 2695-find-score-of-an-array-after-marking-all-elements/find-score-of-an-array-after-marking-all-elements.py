from heapq import *

class Solution:
    def findScore(self, nums: List[int]) -> int:
        minHeap = []
        marked = []
        for i in range(len(nums)):
            heappush(minHeap, (nums[i], i))
            marked.append(False)
        unMarkedCount = len(marked)
        score = 0
        while len(minHeap) > 0:
            minValue = minHeap[0]
            if not marked[minValue[1]]:
                score += minValue[0]
                marked[minValue[1]] = True
                lIndex = max(0, minValue[1]-1)
                rIndex = min(len(nums)-1, minValue[1] + 1)
                marked[lIndex] = True
                marked[rIndex] = True
            else:
                heappop(minHeap)
        return score
        
        
