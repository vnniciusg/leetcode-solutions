impl Solution {
    pub fn min_days(bloom_day: Vec<i32>, m: i32, k: i32) -> i32 {
        
        if k * m > bloom_day.len() as i32 {
            return -1;
        }


        let (mut left, mut right) = (*bloom_day.iter().min().unwrap(), *bloom_day.iter().max().unwrap());
        while left < right {
            let mid = (left + right) / 2;
            if can_make_bouquets(&bloom_day, mid, k, m) {
                right = mid;
            }else {
                left = mid + 1;
            }
        }

        if can_make_bouquets(&bloom_day, left, k, m) {
            return left;
        }

        -1
    }
}

fn can_make_bouquets(bloom_day: &Vec<i32>, day: i32, k: i32, m:i32 ) -> bool {

    let (mut bouquets, mut flowers) = (0, 0);

    for &bloom in bloom_day.iter() {
        
        if bloom <= day {
            flowers += 1;
            if flowers == k {
                bouquets += 1;
                flowers = 0;
            }
        }else {
            flowers = 0;
        }
    }

    bouquets >= m
}
