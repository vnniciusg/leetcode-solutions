impl Solution {
    pub fn find_max_consecutive_ones(nums: Vec<i32>) -> i32 {
        
        let (mut curr_sum, mut max_sum) = (0, 0);

        for &num in nums.iter() {
            if num == 1 {
                curr_sum += 1;
                max_sum = std::cmp::max(max_sum, curr_sum);
            }else {
                curr_sum = 0;
            }
        }

        max_sum
    }
}
