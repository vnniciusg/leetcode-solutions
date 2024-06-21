impl Solution {
    pub fn max_satisfied(customers: Vec<i32>, grumpy: Vec<i32>, minutes: i32) -> i32 {
        
        let lenght_customers = customers.len();

        let mut satisfied_customers = 0;
        for (idx, &costumer) in customers.iter().enumerate() {
            if grumpy[idx] == 0 {
                satisfied_customers += costumer;
            }
        }

        let mut additional_satisfied = 0;
        for idx in 0..minutes as usize {
            if grumpy[idx] == 1 {
                additional_satisfied += customers[idx];
            }
        }

        let mut max_additional_satisfied = additional_satisfied;
        
        for idx in minutes as usize..lenght_customers {
            if grumpy[idx] == 1 {
                additional_satisfied += customers[idx];
            }

            if grumpy[idx - minutes as usize] == 1 {
                additional_satisfied -= customers[idx - minutes as usize];
            }

            max_additional_satisfied = std::cmp::max(max_additional_satisfied, additional_satisfied);
        }

        max_additional_satisfied + satisfied_customers
    }
}
