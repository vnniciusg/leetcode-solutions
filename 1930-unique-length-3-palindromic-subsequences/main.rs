use std::collections::HashSet;

impl Solution {
    pub fn count_palindromic_subsequence(s: String) -> i32 {
        let letters: HashSet<char> = s.chars().collect();
        let chars: Vec<char> = s.chars().collect();
        let mut ans: i32 = 0;

        for &letter in letters.iter() {

            let i = chars.iter().position(|&c| c == letter);
            let j = chars.iter().rposition(|&c| c == letter);

            if let(Some(first), Some(last)) = (i, j) {
                if first == last {
                    continue;
                }

                let between: HashSet<char> = chars[first + 1..last].iter().copied().collect();
                ans += between.len() as i32;
            }
        }
        
        return ans;
    }
}
