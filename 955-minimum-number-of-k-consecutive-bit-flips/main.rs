impl Solution {
    pub fn min_k_bit_flips(nums: Vec<i32>, k: i32) -> i32 {
        
        let mut n = nums.len();
        let (mut flips, mut flips_effect) = (0, 0);

        let mut is_flipped = vec![0; n];

        for idx in 0..n {

            if idx >= k as usize {
                flips_effect ^= is_flipped[idx - k as usize];
            }

            if flips_effect == nums[idx] {

                if idx + k as usize > n {
                    return -1;
                }

                flips += 1;
                flips_effect ^= 1;
                is_flipped[idx] = 1
            }

        }

        flips
    }
}
