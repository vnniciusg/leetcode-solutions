impl Solution {
    pub fn can_place_flowers(flowerbed: Vec<i32>, n: i32) -> bool {

        let mut extended_flowerbed = Vec::with_capacity(flowerbed.len() + 2);
        extended_flowerbed.push(0);
        extended_flowerbed.extend_from_slice(&flowerbed);
        extended_flowerbed.push(0);

        let mut planted_flowers = 0;


        for idx in 1..extended_flowerbed.len() - 1 {
            if extended_flowerbed[idx - 1] == 0 && extended_flowerbed[idx] == 0 && extended_flowerbed[idx + 1] == 0 {
                extended_flowerbed[idx] = 1;
                planted_flowers += 1;
            }
        }

        planted_flowers >= n
    }
}
