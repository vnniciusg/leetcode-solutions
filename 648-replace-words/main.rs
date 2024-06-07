use std::collections::HashSet;

impl Solution {
    pub fn replace_words(dictionary: Vec<String>, sentence: String) -> String {

        let mut root_set: HashSet<String> = HashSet::new();
        for word in dictionary {
            root_set.insert(word);
        }

        let mut sentence_list: Vec<String> = sentence.split_whitespace()
                                                     .map(|s| s.to_string())
                                                     .collect();

        let mut ans: Vec<String> = Vec::new();
        for word in sentence_list {
            ans.push(find_root(word, &root_set));
        }

        ans.join(" ")
    }
}


fn find_root(word: String, root_set: &HashSet<String>) -> String {

    for i in 1..word.len() {
        let prefix: &str = &word[0..i];
        if root_set.contains(prefix) {
            return prefix.to_string();
        }

    }
    return word
}
