class Solution:

    def parseString(self, s, cur):
        res = ""
        # print(cur)
        while cur[0] < len(s) and s[cur[0]] != ']':
            if s[cur[0]] <= 'z' and s[cur[0]] >= 'a':
                res += s[cur[0]]
                cur[0] += 1
            elif s[cur[0]] >= '0' and s[cur[0]] <= '9':
                extractedNum = ''
                while s[cur[0]] >= '0' and s[cur[0]] <= '9':
                    extractedNum += s[cur[0]]
                    cur[0] += 1
                cur[0] += 1
                parsed_string = self.parseString(s, cur)
                cur[0] += 1
                k = int(extractedNum)
                res += ''.join([parsed_string] * k)
        return res


    def decodeString(self, s: str) -> str:
        return self.parseString(s, [0])

        