impl Solution {
    pub fn reverse_words(s: String) -> String {
        
        let mut l: Vec<&str> = s.split_whitespace().collect();
        let n = l.len();

        for i in 0..(n / 2) {
            l.swap(i, n - i - 1);
        }

        l.join(" ").to_string()
    }
}
