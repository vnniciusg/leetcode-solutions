impl Solution {
    pub fn pivot_index(nums: Vec<i32>) -> i32 {
        
        let mut left_sum: i32 = 0;
        let mut right_sum: i32 = nums.iter().sum();

        for (idx, &num) in nums.iter().enumerate() {
            
            right_sum -= num;

            if right_sum == left_sum {
                return idx as  i32;
            }

            left_sum += num;
        }

        -1
    }
}
