impl Solution {
    pub fn special_array(nums: Vec<i32>) -> i32 {

        let n = nums.len();

        for x in 1..=n{
            let count = nums.iter().filter(|&&num| num >= x as i32).count();
            if count == x {
                return x as i32;
            }
        }
    -1
    }
}
