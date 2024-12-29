class Solution:
        def countGoodArrays(self, n: int, m: int, k: int) -> int:
            mod = 10 ** 9 + 7
            noMatchingAdjacentElements = m * pow(m - 1, n - k - 1, mod)
            return noMatchingAdjacentElements * comb(n - 1, k) % mod
        