impl Solution {
    pub fn merge_alternately(word1: String, word2: String) -> String {
        
        let min_length = word1.len().min(word2.len());
        
        let mut ans = String::new();
        for (c1, c2) in word1.chars().zip(word2.chars()) {
            ans.push(c1);
            ans.push(c2);
        }

        ans.push_str(&word1[min_length..]);
        ans.push_str(&word2[min_length..]);

        ans
    }
}
