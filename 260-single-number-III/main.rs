use std::collections::HashMap;

impl Solution {
    pub fn single_number(nums: Vec<i32>) -> Vec<i32> {
        
        let mut count_map: HashMap<i32, i32> = HashMap::new();

        for &num in &nums {
            *count_map.entry(num).or_insert(0) += 1;
        }

        let mut res: Vec<i32> = Vec::new();
        for (&k, &v) in count_map.iter() {
            if v == 1 {
                res.push(k);
            }
        }

        res
    }
}
