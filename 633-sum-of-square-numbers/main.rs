impl Solution {
    pub fn judge_square_sum(c: i32) -> bool {
        
        let mut left: i64 = 0;
        let mut right = (c as f64).sqrt() as i64;


        while left <= right {
            let current_sum = left * left + right * right;
            if current_sum == c as i64 {
                return true;
            }

            if current_sum < c as i64{
                left += 1;
            }else {
                right -= 1;
            }
        }

        false
    }
}
