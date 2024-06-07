impl Solution {
    pub fn smallest_distance_pair(nums: Vec<i32>, k: i32) -> i32 {
        
        let mut sorted_nums: Vec<i32> = nums.clone();
        sorted_nums.sort();

        let n: usize = sorted_nums.len();

        let mut left: i32 = 0;
        let mut right: i32 = sorted_nums[n-1] - sorted_nums[0];

        while left < right {
            let mid = left + (right - left) / 2;
            if count_pairs(&sorted_nums, mid as i32, n) < k {
                left = mid + 1
            }else {
                right = mid
            }
        }

        left as i32
    }
}

fn count_pairs(nums: &[i32], mid: i32, n: usize) -> i32 {
    let mut count = 0;
    let mut left = 0;

    for right in 0..n {
        while nums[right] - nums[left] > mid {
            left += 1;
        } 
        count += (right - left) as i32
    }

    count
}
