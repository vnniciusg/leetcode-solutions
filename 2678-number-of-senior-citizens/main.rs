impl Solution {
    pub fn count_seniors(details: Vec<String>) -> i32 {
        
        let mut count = 0;

        for detail in details.iter() {
            if let Some(age_str) = detail.get(11..13) {
                if let Ok(age) = age_str.parse::<i32>() {
                    if age > 60 {
                        count += 1;
                    }
                }
            }
        }
        count
    }
}
