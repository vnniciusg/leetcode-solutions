use std::collections::HashMap;

impl Solution {
    pub fn check_subarray_sum(nums: Vec<i32>, k: i32) -> bool {
        
        if nums.len() < 2 {
            return false;
        }


        let mut remainder_map: HashMap<i32, i32> = HashMap::new();
        remainder_map.insert(0, -1);

        let mut cumulative_sum: i32 = 0;

        for i in 0..nums.len() {
            cumulative_sum += nums[i];
            let remainder: i32 = cumulative_sum % k;

            if let Some(&prev_index) = remainder_map.get(&remainder) {
                if i as i32 - prev_index > 1 {
                    return true;
                }
            }else {
                remainder_map.insert(remainder, i as i32);
            }
        }

        false

    }
}
