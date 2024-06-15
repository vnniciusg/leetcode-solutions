class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        
        if not grid:
            return 0

        rows, cols = len(grid), len(grid[0])
        visited = set()
        islands = 0

        def dfs(row: int, col: int) -> None:

            if (row, col) in visited or row >= rows or col >= cols or col < 0 or row < 0 or grid[row][col] == "0":
                return
            
            visited.add((row, col))

            directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]
            for (x, y) in directions:
                dfs(row + x, col + y)

        for row in range(rows):
            for col in range(cols):
                if grid[row][col] == "1" and (row, col) not in visited:
                    dfs(row, col)
                    islands += 1
        

        return islands

