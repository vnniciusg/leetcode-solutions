impl Solution {
    pub fn reorder_spaces(text: String) -> String {
        
        if text.is_empty() || !text.contains(' ') {
            return text.to_string();
        }

        let count_spaces =  text.chars().filter(|&c| c == ' ').count();
        let whitout_space_list: Vec<&str> =  text.split_whitespace().collect();
        let m = whitout_space_list.len();

        if m == 1 {
            return whitout_space_list[0].to_string() + &" ".repeat(count_spaces);
        }

        let div_spaces = count_spaces / (m - 1);
        let remaining_spaces = count_spaces % (m -1);

        let mut ans: String = String::new();

        for (i, word) in whitout_space_list.iter().take(m - 1).enumerate() {
            ans += word;
            ans += &" ".repeat(div_spaces);
        }

        ans += whitout_space_list[m - 1]; 
        ans +=  &" ".repeat(remaining_spaces);

        ans
    }
}
