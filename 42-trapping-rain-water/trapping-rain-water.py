class Solution:
    def trap(self, height: List[int]) -> int:
        left, right = 0, len(height) - 1
        maxLeft, maxRight = 0, 0
        totalUnits = 0
        while left <= right:
            if maxLeft < maxRight:
                totalUnits += max(0, maxLeft - height[left])
                maxLeft = max(maxLeft, height[left])
                left+=1
            else:
                totalUnits += max(0, maxRight - height[right])
                maxRight = max(maxRight, height[right])
                right -= 1
        return totalUnits


        