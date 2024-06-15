impl Solution {
    pub fn kids_with_candies(candies: Vec<i32>, extra_candies: i32) -> Vec<bool> {
        
        let mut ans: Vec<bool> = Vec::with_capacity(candies.len());
        let max = match candies.iter().max() {
            Some(&max_value) => max_value,
            None => return ans,
        };

        for &candie in candies.iter() {
            ans.push(candie + extra_candies >= max);
        };

        ans
    }
}
