use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        
        let mut _map = HashMap::new();

        for (idx, &num) in nums.iter().enumerate() {
            if let Some(&value) = _map.get(&num) {
                return vec![idx as i32, value];
            }

            _map.insert(target - num, idx as i32);
        }

        vec![]
    }
}
