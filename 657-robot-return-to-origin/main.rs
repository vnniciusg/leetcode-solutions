impl Solution {
    pub fn judge_circle(moves: String) -> bool {
        
        let mut positions = [0, 0];

        for mov in moves.chars() {
            match mov {
                'U' => positions[0] += 1,
                'D' => positions[0] -= 1,
                'L' => positions[1] -= 1,
                'R' => positions[1] += 1,
                _ => (),
            }
        }

        positions == [0, 0]
    }
}
