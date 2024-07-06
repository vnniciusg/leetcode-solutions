use std::collections::HashSet;

impl Solution {
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {

        let mut nums_set: HashSet<i32> = nums.into_iter().collect();
        let mut longest = 0;

        for &num in &nums_set {
            if !nums_set.contains(&(num - 1)) {
                let mut length = 1;
                while nums_set.contains(&(num + length)) {
                    length += 1;
                }

                longest = longest.max(length);
            }
        }

        longest
    }
}
