impl Solution {
    pub fn diagonal_prime(nums: Vec<Vec<i32>>) -> i32 {
        
        if nums.is_empty() {
            return 0;
        }

        let mut prime_numbers: Vec<i32> = Vec::with_capacity(nums.len() * 2);
        for i in 0..nums.len(){
            prime_numbers.push(nums[i][i]);
            prime_numbers.push(nums[i][nums.len() - i - 1]);
        }

        let mut max_prime = 0;
        for num in prime_numbers {
            if is_prime(num) && num > max_prime {
                max_prime = num;
            }
        }

        max_prime
    }
}

fn is_prime(num: i32) -> bool {
    if num <= 1 {
        return false;
    }

    for i in 2..=((num as f64).sqrt() as i32) {
        if num % i == 0 {
            return false;
        }
    }

    true
}
