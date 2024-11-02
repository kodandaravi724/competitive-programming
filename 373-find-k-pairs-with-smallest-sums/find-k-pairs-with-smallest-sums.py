from heapq import heappop, heappush

class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        res = []
        queue = [(nums1[0] + nums2[0], 0, 0)]
        for _ in range(k):
            _, idx1, idx2 = heappop(queue)
            res.append([nums1[idx1], nums2[idx2]])
            if idx2 == 0 and idx1 + 1 < len(nums1):
                heappush(queue, (nums1[idx1+1] + nums2[0], idx1 + 1, 0))

            if idx2 + 1 < len(nums2):
                heappush(queue, (nums1[idx1] + nums2[idx2 + 1], idx1, idx2 + 1))
        return res