impl Solution {
    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {

        let mut merged_arr = nums1.clone();
        merged_arr.extend(nums2);
        merged_arr.sort();

        let mid = merged_arr.len() / 2;

        if merged_arr.len() & 1 != 0 {
            return merged_arr[mid] as f64;
        }

        (merged_arr[mid] + merged_arr[mid-1]) as f64 / 2.0

    }
}