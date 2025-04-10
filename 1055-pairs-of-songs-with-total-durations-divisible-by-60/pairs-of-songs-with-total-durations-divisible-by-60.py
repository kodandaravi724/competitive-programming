from math import factorial


class Solution:
    def numPairsDivisibleBy60(self, time: List[int]) -> int:
        pairs = [0 for _ in range(60)]
        for t in time:
            tr = t % 60
            pairs[tr] += 1
        ret_num = 0
        for j in [0, 30]:
            if pairs[j] >= 2:
                ret_num += (factorial(pairs[j]) // factorial(pairs[j] - 2) // 2)

        for i in range(1, 30):
            ret_num += (pairs[i] * pairs[-i])
        return ret_num