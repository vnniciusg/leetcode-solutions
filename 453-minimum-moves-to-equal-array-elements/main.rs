impl Solution {
    pub fn min_moves(nums: Vec<i32>) -> i32 {
        let sum: i32 = nums.iter().sum();
        sum - (nums.len() as i32 * nums.iter().min().unwrap())
    }
}