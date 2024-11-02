from heapq import *
class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        result: List[List[int]] = []
        maxHeap = []
        for i in range(min(k, len(nums1))):
            for j in range(min(k, len(nums2))):
                if len(maxHeap) < k:
                    heappush(maxHeap, (-(nums1[i]+nums2[j]), nums1[i], nums2[j]))
                else:
                    currentSum = nums1[i] + nums2[j]
                    if currentSum < -maxHeap[0][0]:
                        heappop(maxHeap)
                        heappush(maxHeap, (-(nums1[i]+nums2[j]), nums1[i], nums2[j]))
                    else:
                        break
        for _, a, b in maxHeap:
            result.append([a, b])
        return result
                
        
        