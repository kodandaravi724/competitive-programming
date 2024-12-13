class Solution:
    def powerValue(self, k, mem):
        if k == 1:
            return 0
        if k in mem:
            return mem[k]
        newValue = k/2 if k%2 == 0 else 3*k + 1
        mem[k] = 1 + self.powerValue(newValue, mem)
        return mem[k]
    def getKth(self, lo: int, hi: int, k: int) -> int:
        t = []
        j = lo
        mem = {}
        while j <= hi:
            t.append((self.powerValue(j, mem), j))
            j+=1
        t.sort()
        return t[k-1][1]
            
        