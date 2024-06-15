use std::collections::HashMap;

impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let mut roman_dict: HashMap<char, i32> = HashMap::new();
        roman_dict.insert('I', 1);
        roman_dict.insert('V', 5);
        roman_dict.insert('X', 10);
        roman_dict.insert('L', 50);
        roman_dict.insert('C', 100);
        roman_dict.insert('D', 500);
        roman_dict.insert('M', 1000);

        let mut total = 0;
        let mut prev = 0;

        for ch in s.chars().rev() {
            if let Some(&value) = roman_dict.get(&ch) {
                if value >= prev {
                    total += value;
                } else {
                    total -= value;
                }
                prev = value;
            }
        }

        total
    }
}
