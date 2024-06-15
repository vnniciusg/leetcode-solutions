use std::collections::HashSet;

impl Solution {
    pub fn reverse_vowels(s: String) -> String {

        let vowels: &str = "aeiouAEIOU";
        let vowel_set: HashSet<char> = vowels.chars().collect();

        let mut chars: Vec<char> = s.chars().collect();
        let (mut left, mut right) = (0, chars.len() - 1);

        while left < right {
            if vowel_set.contains(&chars[left]) && vowel_set.contains(&chars[right]) {
                chars.swap(left, right);
                left += 1;
                right -= 1;
            }
            if !vowel_set.contains(&chars[left]) {
                left += 1;
            }
            if !vowel_set.contains(&chars[right]) {
                right -= 1;
            }
        }

        chars.iter().collect()
    }
}
