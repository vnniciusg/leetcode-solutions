use std::collections::HashMap;

impl Solution {
    pub fn num_jewels_in_stones(jewels: String, stones: String) -> i32 {
        
        let mut stones_map: HashMap<char, i32> = HashMap::new();

        for stone in stones.chars() {
            *stones_map.entry(stone).or_insert(0) += 1;
        }

        let mut count: i32 = 0;
        for jewel in jewels.chars() {
            if let Some(&stone_count) = stones_map.get(&jewel) {
                count += stone_count;
            }
        }

        count
    }   
}
