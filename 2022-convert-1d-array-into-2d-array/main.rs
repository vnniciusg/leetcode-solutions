impl Solution {
    pub fn construct2_d_array(original: Vec<i32>, m: i32, n: i32) -> Vec<Vec<i32>> {
        
        if original.len() != (m * n) as usize {
            return Vec::new();
        }

        let mut ans: Vec<Vec<i32>> = Vec::new();
        for i in 0..m {
            let row: Vec<i32> = original[(i*n) as usize..((i+1) * n) as usize].to_vec();
            ans.push(row);
        }

        ans
    }
}
