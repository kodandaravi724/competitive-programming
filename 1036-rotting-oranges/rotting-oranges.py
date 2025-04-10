class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        rottenOranges = []
        i = 0
        m = len(grid)
        n = len(grid[0])
        while i < m:
            j = 0
            while j < n:
                if grid[i][j] == 2:
                    rottenOranges.append((i, j))
                j += 1
            i += 1
        dirs = [[1, 0], [-1, 0], [0, 1], [0, -1]]
        minutes = 0
        while rottenOranges:
            index = 0
            num = len(rottenOranges)
            # print(rottenOranges, minutes)
            freshOrange = False
            while index < num:
                x, y = rottenOranges.pop(0)
                for dx, dy in dirs:
                    nx, ny = x + dx, y + dy
                    while (nx >= 0 and nx < m) and (ny >= 0 and ny < n) and grid[nx][ny] == 1:
                        freshOrange = True
                        grid[nx][ny] = 2
                        # print(nx, ny)
                        rottenOranges.append((nx, ny))
                index += 1
            if freshOrange:
                minutes += 1
        i = 0
        while i < m:
            j = 0
            while j < n:
                if grid[i][j] == 1:
                    return -1
                j += 1
            i += 1
        return minutes
