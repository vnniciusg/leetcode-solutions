impl Solution {
    pub fn equal_substring(s: String, t: String, max_cost: i32) -> i32 {
        
        let mut max_len = 0;
        let mut cost = 0;
        let mut left = 0;
        let mut right = 0;

        let s = s.as_bytes();
        let t = t.as_bytes();

        while right < s.len() {
            cost += (s[right] as i32 - t[right] as i32).abs();
            right += 1;

            while cost > max_cost {
                cost -= (s[left] as i32 - t[left] as i32).abs();
                left += 1;
            }

            max_len = max_len.max(right - left);
        }

        max_len as i32
    }


}