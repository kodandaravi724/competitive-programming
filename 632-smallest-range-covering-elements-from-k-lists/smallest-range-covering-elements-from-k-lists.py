from heapq import *

class Solution:
    def smallestRange(self, nums: List[List[int]]) -> List[int]:
        minHeap = []
        curMaxNumber = -1
        range_start = float('-inf')
        range_end = float('inf')
        for i in range(len(nums)):
            heappush(minHeap, (nums[i][0], 0, nums[i]))
            curMaxNumber = max(curMaxNumber, nums[i][0])
        n = len(nums)
        while len(minHeap) == n:
            value, index, arr = heappop(minHeap)
            if range_end - range_start > curMaxNumber - value:
                range_start = value
                range_end = curMaxNumber
            if index+1 < len(arr):
                heappush(minHeap, (arr[index+1], index+1, arr))
                curMaxNumber = max(curMaxNumber, arr[index+1])
        return [range_start, range_end]

        