use std::collections::HashMap;

impl Solution {
    pub fn max_operations(nums: Vec<i32>, k: i32) -> i32 {
        
        let mut count_map: HashMap<i32, i32> = HashMap::new();
        let mut count_pairs = 0;


        for &num in &nums {
            let complement = k - num;
            match count_map.get_mut(&complement) {
                Some(count) if *count > 0 => {
                    count_pairs += 1;
                    *count -= 1;
                },
                _ => {
                    *count_map.entry(num).or_insert(0) += 1;
                }
            }
        }

        count_pairs
    }
}
