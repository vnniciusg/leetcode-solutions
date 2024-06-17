impl Solution {
    pub fn maximum_difference(nums: Vec<i32>) -> i32 {
        
        if nums.is_empty() {
            return 0;
        }

        let mut max_difference = -1;
        let mut min_value = nums[0];

        for (idx, &num) in nums.iter().enumerate() {
            if num > min_value {
                max_difference = std::cmp::max(max_difference, num - min_value);
            }else {
                min_value = num;
            }
        } 

        max_difference
    }
}
