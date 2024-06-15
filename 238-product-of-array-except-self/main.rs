impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut ans = vec![1; n];
        
        let mut prefix = 1;
        for i in 0..n {
            ans[i] *= prefix;
            prefix *= nums[i];
        }
        
        let mut suffix = 1;
        for i in (0..n).rev() {
            ans[i] *= suffix;
            suffix *= nums[i];
        }
        
        ans
    }
}
