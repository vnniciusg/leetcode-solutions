class Solution:
    def highestPeak(self, isWater: List[List[int]]) -> List[List[int]]:
        
        m, n = len(isWater), len(isWater[0])
        height = [[-1 for _ in range(n)] for _ in range(m)]
        q = deque()

        for i in range(m):
            for j in range(n):
                if isWater[i][j] == 1:
                    height[i][j] = 0
                    q.append((i, j))

        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]

        while q:
            i, j = q.popleft()
            for dx, dy in directions:
                ni, nj = i + dx, j + dy

                if 0 <= ni < m and 0 <= nj < n and height[ni][nj] == -1:
                    height[ni][nj] = height[i][j] + 1
                    q.append((ni, nj))
        

        return height
