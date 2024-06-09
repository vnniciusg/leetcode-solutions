use std::collections::HashMap;

impl Solution {
    pub fn subarrays_div_by_k(nums: Vec<i32>, k: i32) -> i32 {
        
        let mut count = 0;
        let mut prefix_sum = 0;

        let mut mod_count: HashMap<i32, i32> = HashMap::new();
        mod_count.insert(0, 1);

        for num in nums.iter() {
            prefix_sum += num;
            let mut curr_mod = prefix_sum % k;

            if curr_mod < 0 {
                curr_mod += k;
            }

            if let Some(&value) = mod_count.get(&curr_mod) {
                count += value;
            }

            *mod_count.entry(curr_mod).or_insert(0) += 1;
        }

        count
    }
}
