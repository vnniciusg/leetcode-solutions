use std::collections::HashMap;

impl Solution {
    pub fn longest_palindrome(s: String) -> i32 {
        
        let mut freq = HashMap::new();

        for ch in s.chars(){
            *freq.entry(ch).or_insert(0) += 1;
        }

    
        let mut length = 0;
        let mut odd_found = false;

        for &value in freq.values() {
            if value & 1 == 0 {
                length += value;
            }else {
                length += value - 1;
                odd_found = true;
            }
        }

        if odd_found {
            length += 1
        }

        length
    }
}
