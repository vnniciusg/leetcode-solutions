impl Solution {
    pub fn largest_palindromic(num: String) -> String {
        
        let mut cnt = vec![0; 10];
        for c in num.chars() {
            cnt[c.to_digit(10).unwrap() as usize] += 1;
        }

        let mut ans = String::new();
        for i in (0..=9).rev() {
            if cnt[i] % 2 == 1 {
                ans.push_str(&i.to_string());
                cnt[i] -= 1;
                break;
            }
        }

        for i in 0..10 {
            if cnt[i] > 0 {
                cnt[i] /= 2;
                let s = std::iter::repeat(i.to_string()).take(cnt[i]).collect::<String>();
                ans = s.clone() + &ans + &s;
            }
        }
        
        ans = ans.trim_matches('0').to_string();
        if ans.is_empty() {
           return "0".to_string()
        } 

        ans

    }
}


