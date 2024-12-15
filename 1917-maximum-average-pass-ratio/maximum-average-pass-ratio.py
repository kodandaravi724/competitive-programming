from heapq import *

class Solution:
    def calculate_gain(self, a, b):
        return ((a+1)/(b+1)) - (a/b)

    def maxAverageRatio(self, classes: List[List[int]], extraStudents: int) -> float:
        maxHeap = []

        n = len(classes)

        for _class in classes:
            heappush(maxHeap, (-self.calculate_gain(_class[0], _class[1]), _class[0], _class[1]))

        while extraStudents > 0:
            gain, a, b = heappop(maxHeap)
            heappush(maxHeap, (-self.calculate_gain(a+1, b+1), a+1, b+1))
            extraStudents-=1
        res = 0
        for a, b, c in maxHeap:
            res += (b/c)
        return res/n
        


        

        