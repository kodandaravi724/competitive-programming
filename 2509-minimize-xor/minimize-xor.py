class Solution:

    def count_ones(self, x):
        res = 0
        while x != 0:
            if x & 1 == 1:
                res += 1
            x = x >> 1
        return res

    def minimizeXor(self, num1: int, num2: int) -> int:
        t1 = num1
        t2 = num2
        n = self.count_ones(num2)
        t = t1 >> n
        ones_1 = self.count_ones(num1)
        print(ones_1, n)
        if t == 0:
            return (1 << n) - 1
        else:
            if ones_1 > n:
                d = (ones_1 - n)
                while d > 0:
                    t1 = (t1-1)&t1
                    d-=1
                return t1
            else:
                excess = n - ones_1
                t = 1
                shift = 0
                res = 0
                print(excess, "excess")
                while num1 != 0 and excess > 0:
                    while num1 & 1 != 0:
                        num1 = num1 >> 1
                        # t = 1 << shift
                        shift+=1
                    print(t)
                    num1 = num1 >> 1
                    shift += 1
                    t = 1 << (shift-1)
                    res = res | t
                    excess -= 1
                return t1 | res


        
        