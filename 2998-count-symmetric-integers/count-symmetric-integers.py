class Solution:
    def countSymmetricIntegers(self, low: int, high: int) -> int:
        start = max(11, low)
        numSymmetricIntegers = 0
        while start <= high:
            if start >= 11 and start < 100:
                if (start % 11) == 0:
                    numSymmetricIntegers += 1
            elif (start//1000 == 0):
                start = 1001
                continue
            else:
                l = start//1000 + (start//100) % 10
                r = start%10 + (start%100)//10
                # print(l, r, start)
                if l == r:
                    numSymmetricIntegers += 1
            start += 1
        return numSymmetricIntegers