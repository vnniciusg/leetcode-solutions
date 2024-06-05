use std::collections::HashMap;

impl Solution {
    pub fn common_chars(words: Vec<String>) -> Vec<String> {
        if words.is_empty() {
            return Vec::new();
        }

        if words.len() == 1 {
            return words[0].chars().map(|c| c.to_string()).collect();
        }

        let mut common_freq: HashMap<char, i32> = HashMap::new();
        for char in words[0].chars() {
            *common_freq.entry(char).or_insert(0) += 1;
        }

        for word in words.iter().skip(1) {
            let mut word_freq: HashMap<char, i32> = HashMap::new();
            for char in word.chars(){
                *word_freq.entry(char).or_insert(0) += 1;
            }

            for key in common_freq.keys().cloned().collect::<Vec<char>>() {
                if let Some(&count) = word_freq.get(&key){
                    *common_freq.get_mut(&key).unwrap() = std::cmp::min(common_freq[&key], count);
                }else {
                    common_freq.remove(&key);
                }
            }
        }

        let mut ans: Vec<String> = Vec::new();
        for (char, freq) in common_freq {
            for _ in 0..freq {
                ans.push(char.to_string());
            }
        }

        ans

    }
}