impl Solution {
    pub fn score_of_string(s: String) -> i32 {
        
        let n = s.len();
        let mut score = 0;

        for i in 0..(n-1){
            let current_digit = s.chars().nth(i).unwrap() as i32 - '0' as i32;
            let next_digit = s.chars().nth(i+1).unwrap() as i32 - '0' as i32;
            score += (current_digit - next_digit).abs();
        }

        score
    }

}
