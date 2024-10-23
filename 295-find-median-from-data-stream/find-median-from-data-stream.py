from heapq import *

class MedianFinder:

    def __init__(self):
        self.minHeap = []
        self.maxHeap = []
        

    def addNum(self, num: int) -> None:
        if len(self.maxHeap) == 0:
            heappush(self.maxHeap, -num)
        else:
            if len(self.maxHeap) == len(self.minHeap) + 1:
                if -self.maxHeap[0] > num:
                    heappush(self.maxHeap, -num)
                    heappush(self.minHeap, -heappop(self.maxHeap))
                else:
                    heappush(self.minHeap, num)
            else:
                if(self.minHeap[0] < num):
                    heappush(self.minHeap, num)
                    heappush(self.maxHeap, -heappop(self.minHeap))
                else:
                    heappush(self.maxHeap, -num)

    def findMedian(self) -> float:
        # print(self.minHeap , self.maxHeap)
        if len(self.minHeap) != len(self.maxHeap):
            return -self.maxHeap[0]
        else:
            return (self.minHeap[0] - self.maxHeap[0]) / 2
        


# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()