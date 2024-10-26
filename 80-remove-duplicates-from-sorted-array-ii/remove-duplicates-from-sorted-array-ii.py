class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        cur = 0
        start = 0
        cnt = 0
        prev = None
        n = len(nums)
        while start < len(nums):
            if prev is not None:
                if prev == nums[start]:
                    if cnt < 2:
                        prev = nums[start]
                        nums[start], nums[cur] = nums[cur], nums[start]
                        cur+=1
                        cnt+=1
                elif prev != nums[start]:
                    prev = nums[start]
                    cnt = 1
                    nums[start], nums[cur] = nums[cur], nums[start]
                    cur+=1
                start += 1
            else:
                prev = nums[start]
                nums[start], nums[cur] = nums[cur], nums[start]
                cur+=1
                cnt+=1
                start+=1
        return cur

        