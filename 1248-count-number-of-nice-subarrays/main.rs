use std::collections::HashMap;

impl Solution {
    pub fn number_of_subarrays(nums: Vec<i32>, k: i32) -> i32 {
        
        let mut map: HashMap<i32, i32> = HashMap::new();
        map.insert(0, 1);

        let (mut count, mut odd_count) = (0, 0);

        for &num in nums.iter() {
            
            if num & 1 != 0 {
                odd_count += 1;
            }
            
            if let Some(value) = map.get(&(odd_count - k)) {
                count += *value;
            }

            *map.entry(odd_count).or_insert(0) += 1;
        }

        count
    }
}
