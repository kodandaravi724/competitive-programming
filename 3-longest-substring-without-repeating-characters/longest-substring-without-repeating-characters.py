class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        windowStart, windowEnd = 0, 0
        frequencyMap = {}
        maxLength = float('-inf')
        n = len(s)
        while windowEnd < n:
            frequencyMap[s[windowEnd]] = frequencyMap.get(s[windowEnd], 0) + 1
            while frequencyMap[s[windowEnd]] > 1:
                frequencyMap[s[windowStart]] -= 1
                if frequencyMap[s[windowStart]] == 0:
                    del frequencyMap[s[windowStart]]
                windowStart += 1
            maxLength = max(windowEnd - windowStart + 1, maxLength)
            windowEnd += 1
        return maxLength if maxLength != float('-inf') else 0
