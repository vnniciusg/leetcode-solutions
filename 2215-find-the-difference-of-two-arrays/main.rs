use std::collections::HashSet;

impl Solution {
    pub fn find_difference(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<Vec<i32>> {

        let set1: HashSet<i32> = nums1.into_iter().collect();
        let set2: HashSet<i32> = nums2.into_iter().collect();


        let diff1: Vec<i32> = set1.difference(&set2).cloned().collect();
        let diff2: Vec<i32> = set2.difference(&set1).cloned().collect();

        vec![diff1, diff2]
    }
}
