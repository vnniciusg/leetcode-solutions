use std::collections::HashMap;


impl Solution {
    pub fn subarray_sum(nums: Vec<i32>, k: i32) -> i32 {
        
        let mut ans = 0;
        let mut prefix_sum = 0;

        let mut count: HashMap<i32, i32> = HashMap::new();  
        count.insert(0, 1);

        for num in nums.iter() {
            prefix_sum += num;
            if let Some(&c) = count.get(&(prefix_sum - k)) {
                ans += c;
            }

            *count.entry(prefix_sum).or_insert(0) += 1;
        }


        ans
        
    }
}
