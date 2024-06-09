use std::collections::HashMap;

impl Solution {
    pub fn count_bad_pairs(nums: Vec<i32>) -> i64 {
        
        let n = nums.len();
        let total_pairs: i64 = (n * (n-1) / 2) as i64;   
        let mut good_pairs: i64 = 0;

        let mut diff_count: HashMap<i32, i32> = HashMap::new();

        for i in 0..n {
            let diff = nums[i] - i as i32;
            if let Some(&count) = diff_count.get(&diff) {
                good_pairs += count as i64;
            }
            *diff_count.entry(diff).or_insert(0) += 1;
        }

        total_pairs - good_pairs
    }
}
