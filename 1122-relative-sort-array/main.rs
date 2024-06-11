use std::collections::HashMap;
use std::iter::repeat;

impl Solution {
    pub fn relative_sort_array(arr1: Vec<i32>, arr2: Vec<i32>) -> Vec<i32> {
        
        let mut order_map: HashMap<i32, usize> = HashMap::new();
        for (index, &value) in arr2.iter().enumerate() {
            order_map.insert(value, index);   
        }

        let mut count_map: HashMap<i32, i32> = HashMap::new();
        for &num in arr1.iter() {
            *count_map.entry(num).or_insert(0) += 1;
        }

        let mut ans: Vec<i32> = Vec::new();
        for &num in arr2.iter() {
            if let Some(&count) = count_map.get(&num) {
                ans.extend(repeat(num).take(count as usize));
                count_map.remove(&num);
            }
        }

        let mut remaining: Vec<i32> = Vec::new();
        for (&num, &count) in count_map.iter() {
           remaining.extend(repeat(num).take(count as usize));
        }
        remaining.sort();

        ans.extend(remaining);

        ans
    }
}
