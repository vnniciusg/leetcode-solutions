impl Solution {
    pub fn fizz_buzz(n: i32) -> Vec<String> {
        let mut ans: Vec<String> = Vec::new();

        for i in 1..=n {
            match (i % 3, i % 5) {
                (0, 0) => ans.push("FizzBuzz".to_string()),
                (0, _) => ans.push("Fizz".to_string()),
                (_, 0) => ans.push("Buzz".to_string()),
                _ => ans.push(i.to_string()),
            }
        }

        ans
    }
}

