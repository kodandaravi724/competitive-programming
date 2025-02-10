from collections import deque

class Solution:
    def clearDigits(self, s: str) -> str:
        stack = deque([])
        for c in s:
            if c >= '0' and c <= '9':
                stack.pop()
            else:
                stack.append(c)
        return ''.join(stack)

        