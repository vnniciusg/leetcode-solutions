func numIslands(grid [][]byte) int {
    
    rows, cols := len(grid), len(grid[0])

    if rows == 0 {
        return 0
    }

    visited := make(map[[2]int]bool)
    islands := 0

    var dfs func(row, col int)
    dfs = func (row, col int) {
        if visited[[2]int{row, col}] || row >= rows || row < 0 || col >= cols || col < 0 || grid[row][col] == '0' {
            return
        }

        if _, seen := visited[[2]int{row, col}]; seen {
            return
        }

        visited[[2]int{row, col}] = true
        directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
        for _, dir := range directions {
            r, c := row+dir[0], col+dir[1]
            dfs(r, c)
        }
    }

    for row := range(rows) {
        for col := range(cols) {
            if grid[row][col] == '1' && !visited[[2]int{row, col}] {
                dfs(row, col)
                islands++
            }
        }
    }

    return islands

}
