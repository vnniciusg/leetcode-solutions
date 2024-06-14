impl Solution {
    pub fn min_increment_for_unique(nums: Vec<i32>) -> i32 {
        
        let mut sorted_nums = nums.clone();
        sorted_nums.sort();

        let mut increments = 0;
        for i in 1..sorted_nums.len() {
            if sorted_nums[i] <= sorted_nums[i - 1] {
                let needed = sorted_nums[i - 1] + 1 - sorted_nums[i];
                sorted_nums[i] += needed;
                increments += needed;
            }
        }

        increments
    }
}
