impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        
        let mut max_amount = 0;
        let (mut i, mut j): (usize, usize) = (0, height.len() - 1);

        while i < j {
            let curr_amount = std::cmp::min(height[i], height[j]) * (j - i) as i32;
            if curr_amount > max_amount {
                max_amount = curr_amount;
            }else if height[i] < height[j] {
                i += 1;
            }else {
                j -= 1;
            }
        }

        max_amount
    }
}
