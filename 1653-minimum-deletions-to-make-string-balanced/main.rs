impl Solution {
    pub fn minimum_deletions(s: String) -> i32 {
        
        let (mut b_count, mut deletions) = (0, 0);

        for chr in s.chars() {
            match chr {
                'b' => b_count += 1,
                _ => deletions = std::cmp::min(deletions + 1, b_count),
            }
        }

        deletions
    }
}
