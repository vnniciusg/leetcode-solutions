use std::collections::HashMap;

impl Solution {
    pub fn is_n_straight_hand(hand: Vec<i32>, group_size: i32) -> bool {

        if &hand.len() % group_size as usize  != 0 {
            return false;
        }

        let mut count = HashMap::new();
        for &card in &hand {
            *count.entry(card).or_insert(0) += 1;
        }

        let mut sorted_hand = hand.clone();
        sorted_hand.sort();

        for &card in &sorted_hand {
            if let Some(&card_count) = count.get(&card) {
                if card_count > 0 {
                    for i in 0..group_size {
                        let current_card = card + i;
                        if let Some(count) = count.get_mut(&current_card) {
                            if *count > 0 {
                                *count -= 1;
                            } else {
                                return false;
                            }
                        } else {
                            return false;
                        }
                    }
                }
            }
        }

        true
    }
}
