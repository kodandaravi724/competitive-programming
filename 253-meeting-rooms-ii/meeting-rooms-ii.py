class Solution:
    def minMeetingRooms(self, intervals: List[List[int]]) -> int:
        maxTime = 10**6
        numMeetings = [0]*(maxTime + 1)
        for interval in intervals:
            numMeetings[interval[0]] += 1
            numMeetings[interval[1]] -= 1
            # print(interval[1]+1)
        minMeetingRooms = float('-inf')
        # print(numMeetings)
        meetings = 0
        for i in range(maxTime+1):
            meetings += numMeetings[i]
            # print(meetings)
            minMeetingRooms = max(minMeetingRooms, meetings)
        return minMeetingRooms