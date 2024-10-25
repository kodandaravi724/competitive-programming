class FreqStack:

    def __init__(self):
        self.stack = []
        self.maxFrequency = -1
        self.frequencyMap = {}
        

    def push(self, val: int) -> None:
        self.frequencyMap[val] = self.frequencyMap.get(val, 0) + 1
        self.maxFrequency = max(self.frequencyMap[val], self.maxFrequency)
        if self.maxFrequency > len(self.stack):
            self.stack.append([])
        self.stack[self.frequencyMap[val] - 1].append(val)

    def pop(self) -> int:
        element = self.stack[self.maxFrequency - 1].pop()
        if len(self.stack[self.maxFrequency - 1]) == 0:
            self.stack.pop()
        self.maxFrequency = len(self.stack)
        self.frequencyMap[element]-=1
        if self.frequencyMap[element] == 0:
            del self.frequencyMap[element]
        return element


# Your FreqStack object will be instantiated and called as such:
# obj = FreqStack()
# obj.push(val)
# param_2 = obj.pop()