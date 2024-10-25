class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        windowStart, windowEnd = 0, 0
        minimalLength = float('inf')
        currentSum = 0
        while windowEnd < len(nums):
            currentSum += nums[windowEnd]
            while currentSum >= target:
                minimalLength = min(minimalLength, windowEnd - windowStart + 1)
                currentSum -= nums[windowStart]
                windowStart+=1
            windowEnd += 1
        return minimalLength if minimalLength != float('inf') else 0
        