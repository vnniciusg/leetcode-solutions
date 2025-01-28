class Solution:
    def findMaxFish(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        visited = [[False] * cols for _  in range(rows)]
        result = 0

        def dfs(row, col):
            if (
                row < 0
                or row >= rows
                or col < 0
                or col >= cols
                or grid[row][col] == 0
                or visited[row][col]
            ):
                return 0
            
            visited[row][col] = True

            directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]

            return grid[row][col] + sum([dfs(row + dx, col + dy) for dx, dy in directions])

        for row in range(rows):
            for col in range(cols):
                if grid[row][col] > 0 and not visited[row][col]:
                    result = max(result, dfs(row, col))


        return result
