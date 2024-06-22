impl Solution {
    pub fn kth_factor(n: i32, k: i32) -> i32 {
        
        let mut factor_list: Vec<i32> = Vec::new();
        for idx in 1..(n+1) as usize {
            if n % idx as i32 == 0 {
                factor_list.push(idx as i32);
            }
        }

        if k - 1 >= factor_list.len() as i32 {
            return -1;
        }

        factor_list[k as usize - 1]
    }
}
