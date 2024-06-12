impl Solution {
    pub fn divide_string(s: String, k: i32, fill: char) -> Vec<String> {
        
        let mut split_string_list: Vec<String> = Vec::new();

        for i in (0..s.len()).step_by(k as usize) {
            let end = std::cmp::min(i +k as usize, s.len());
            let mut substr = String::from(&s[i..end]);

            if substr.len() != k as usize {
                substr.extend(std::iter::repeat(fill).take(k as usize - substr.len()));
            }

            split_string_list.push(substr);
        }

        split_string_list
    }
}
