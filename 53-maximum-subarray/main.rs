impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        
        let mut cur_sum: i32 = 0;
        let mut max_sum: i32 = std::i32::MIN;;

        let n = nums.len();

        for i in 0..n {
            cur_sum +=  nums[i];
            max_sum = max_sum.max(cur_sum);
            if cur_sum < 0 {
                cur_sum = 0
            }

        }

        max_sum
    }
}
