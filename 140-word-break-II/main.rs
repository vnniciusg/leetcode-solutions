impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> Vec<String> {
        let mut word_set = std::collections::HashSet::new();
        for word in word_dict {
            word_set.insert(word);
        }

        let mut sentences = Vec::new();
        let n = s.len() as i32;

        Self::backtrack(0, &mut Vec::new(), &s, n, &word_set, &mut sentences);
        sentences
    }

    fn backtrack(start: i32, path: &mut Vec<String>, s: &String, n: i32, word_set: &std::collections::HashSet<String>, sentences: &mut Vec<String>) {
        if start == n {
            sentences.push(path.join(" "));
            return;
        }

        for i in start..n {
            let word = &s[start as usize..i as usize + 1];
            if word_set.contains(word) {
                path.push(word.to_string());
                Self::backtrack(i + 1, path, s, n, word_set, sentences);
                path.pop();
            }
        }
    }
}