use std::collections::HashSet;

impl Solution {
    pub fn distinct_prime_factors(nums: Vec<i32>) -> i32 {

        let mut distinct_primes = HashSet::new();
        for &num in nums.iter() {
            let mut factors = prime_factors(num);
            for &factor in &factors {
                distinct_primes.insert(factor);
            }
        }

        distinct_primes.len() as i32
        
    }
}

fn prime_factors(mut n: i32) -> HashSet<i32> {
            
    let mut factors = HashSet::new();
    let mut i = 2;
            
    while n > 1 {
        if n % i == 0 {
            factors.insert(i);
            n /= i;
        }else{
            i += 1;
        }
    }

    factors
}
