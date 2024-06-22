use std::collections::HashMap;

impl Solution {
    pub fn num_identical_pairs(nums: Vec<i32>) -> i32 {
        
        let mut count_values: HashMap<i32, i32> = HashMap::new();
        let mut count_good_pairs = 0;

        for &num in nums.iter() {
            if let Some(value) = count_values.get(num) {
                count_good_pairs += value;
            }
            count_values[num]++
        }

        count_good_pairs
    }
}
