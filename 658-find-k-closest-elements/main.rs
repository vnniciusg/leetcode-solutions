impl Solution {
    pub fn find_closest_elements(arr: Vec<i32>, k: i32, x: i32) -> Vec<i32> {
        
        let mut left: usize = 0;
        let mut right = arr.len() - 1;

        while (right - left) as i32 >= k {
            if (arr[left] as i32 - x).abs() > (arr[right] as i32 - x).abs() {
                left += 1;
            }else {
                right -= 1;
            }
        }

        arr[left..right + 1].to_vec()
    }
}
