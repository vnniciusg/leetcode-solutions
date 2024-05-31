use std::collections::HashMap;
use std::collections::HashSet;

impl Solution {
    pub fn unique_occurrences(arr: Vec<i32>) -> bool {
        
        let mut freq: HashMap<i32,i32 > = HashMap::new();

        for &num in &arr {
            *freq.entry(num).or_insert(0) += 1;
        }

        let mut occurrence = HashSet::new();
        for &count in freq.values(){
            if !occurrence.insert(count) {
                return false;
            }
        }
        true
    }
}
