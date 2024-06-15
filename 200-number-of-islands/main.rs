use std::collections::HashSet;

impl Solution {
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        
        if grid.is_empty() {
            return 0
        }

        let rows_length = grid.len();
        let cols_length = grid[0].len();

        let mut visited: HashSet<(usize, usize)> = HashSet::new();   
        let mut islands = 0;

        for row in 0..rows_length {
            for col in 0..cols_length {
                if grid[row][col] == '1' && !visited.contains(&(row, col)) {
                    dfs(row, col, &grid, &mut visited);
                    islands += 1;
                }
            }
        }

        islands
    }
}


fn dfs(row: usize, col: usize, grid: &Vec<Vec<char>>, visited: &mut HashSet<(usize, usize)>) {
        if row < 0 || col < 0 || row >= grid.len() || col >= grid[0].len() || grid[row][col] == '0' || visited.contains(&(row, col)) {
            return;
        }
        visited.insert((row, col));

        let directions: Vec<(i32, i32)> = [(1, 0), (-1, 0), (0, 1), (0, -1)].to_vec();
        for (dr, dc) in &directions {
            let (r, c) = ((row as i32 + dr) as usize, (col as i32 + dc) as usize);
            dfs(r , c, grid, visited);
        }
    }
