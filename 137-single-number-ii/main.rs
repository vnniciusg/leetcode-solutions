impl Solution {
    pub fn single_number(nums: Vec<i32>) -> i32 {
        let mut ones = 0;
        let mut twos = 0;

        for num in nums.iter() {
            ones = (ones ^ num) & !twos;
            twos = (twos ^ num) & !ones;
        }

        ones
    }
}
