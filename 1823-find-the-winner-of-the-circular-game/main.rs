impl Solution {
    pub fn find_the_winner(n: i32, k: i32) -> i32 {
        let mut friends: Vec<i32> = (1..=n).collect();
        let mut idx: usize = 0;

        while friends.len() > 1 {
            idx = (idx + k as usize - 1) % friends.len();
            friends.remove(idx);
        }

        friends[0]
    }
}