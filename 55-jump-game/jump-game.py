class Solution:
    def canJump(self, nums: List[int]) -> bool:
        maxIndex = 0
        for index in range(len(nums)):
            maxIndex = max(maxIndex, index+nums[index])
            if maxIndex == len(nums) - 1:
                return True
            if maxIndex <= index:
                return False
        return True