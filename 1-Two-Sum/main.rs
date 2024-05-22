impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut result = Vec::new();
        let n = nums.len();

        for i in 0..n-1 {
            for j in i+1..n {
                if nums[i] + nums[j] == target {
                    result.push(i as i32);
                    result.push(j as i32);
                }
            }
        }

        result
    }
}
