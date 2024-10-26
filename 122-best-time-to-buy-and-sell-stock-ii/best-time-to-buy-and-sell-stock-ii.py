class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        prefixDif = []
        for i in range(0, len(prices)-1):
            prefixDif.append(prices[i] - prices[i+1])
        return -sum([profit for profit in prefixDif if profit < 0])
        