impl Solution {
    pub fn min_patches(nums: Vec<i32>, n: i32) -> i32 {
        
        let mut sum: i64 = 0;
        let mut cnt = 0;
        let mut i: usize  = 0;

        while sum < n as i64 {
            if i < nums.len() && nums[i] as i64 <= sum + 1 {
                sum += nums[i] as i64;
                i += 1;
            }else {
                cnt += 1;
                sum += sum + 1;
            }
        }

        cnt
    }
}
