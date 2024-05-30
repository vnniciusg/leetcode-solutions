impl Solution {
    pub fn count_triplets(arr: Vec<i32>) -> i32 {
        
        let n: usize = arr.len();
        let mut prefix = vec![0; n+1];

        for i in 0..n {
            prefix[i+1] = prefix[i] ^ arr[i];
        }

        let mut count = 0;
        for j in 0..n {
            for k in (j+1)..n {
                if prefix[j] == prefix[k+1] {
                    count += (k-j) as i32;
                }
            }
        }

        count

    }
}
