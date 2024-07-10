impl Solution {
    pub fn min_operations(logs: Vec<String>) -> i32 {
        let mut depth = 0;

        for log in logs.iter() {
            match log.as_str() {
                "../" => {
                    if depth > 0 {
                        depth -= 1;
                    }
                },
                "./" => continue,
                _ => depth += 1,
            }
        }

        depth
    }
}
