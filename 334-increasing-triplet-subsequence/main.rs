impl Solution {
    pub fn increasing_triplet(nums: Vec<i32>) -> bool {
        
        if nums.len() < 3 {
            return false;
        }

        let mut first = i32::max_value();
        let mut second = i32::max_value();

        for &num in nums.iter() {
            if num <= first {
                first = num;
            }else if num <= second {
                second = num;
            }else {
                return true;
            }
        }

        false
    }
}
