from collections import defaultdict

class ATM:

    def __init__(self):
        self.denominations = [20, 50, 100, 200, 500]
        self.freq = defaultdict(int)
        

    def deposit(self, banknotesCount: List[int]) -> None:
        for i, bankNotes in enumerate(banknotesCount):
            self.freq[self.denominations[i]] += bankNotes
        

    def withdraw(self, amount: int) -> List[int]:
        currentState = self.freq.copy()
        i = 4
        res = [0]*5
        while i >= 0 and amount > 0:
            curDenomination = self.denominations[i]
            if self.freq[curDenomination] > 0 and curDenomination <= amount:
                numDenominations = min(self.freq[curDenomination], (amount // curDenomination))
                self.freq[curDenomination] -= numDenominations
                res[i] = numDenominations
                amount -= numDenominations*curDenomination
            i -= 1
        if amount > 0:
            self.freq = currentState.copy()
            return [-1]
        return res
        
        


# Your ATM object will be instantiated and called as such:
# obj = ATM()
# obj.deposit(banknotesCount)
# param_2 = obj.withdraw(amount)