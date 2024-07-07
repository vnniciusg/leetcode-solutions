impl Solution {
    pub fn num_water_bottles(num_bottles: i32, num_exchange: i32) -> i32 {
        
        let (mut total_drunk, mut empty_bottles) = (num_bottles, num_bottles);

        while empty_bottles >= num_exchange {

            let new_full_bottles = empty_bottles / num_exchange;
            total_drunk += new_full_bottles;

            empty_bottles = empty_bottles % num_exchange + new_full_bottles;
        }

        total_drunk
    }
}
