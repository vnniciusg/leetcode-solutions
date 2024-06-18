impl Solution {
    pub fn max_profit_assignment(difficulty: Vec<i32>, profit: Vec<i32>, worker: Vec<i32>) -> i32 {
        
        let mut n = difficulty.len();
        let mut jobs: Vec<(i32, i32)> = Vec::with_capacity(n);
        for i in 0..n {
            jobs.push((difficulty[i], profit[i]));
        }
        jobs.sort();

        let mut sorted_worker = worker.clone();
        sorted_worker.sort();

        let (mut max_profit, mut best_profit) = (0, 0);
        let mut job_index: usize = 0;
        n = sorted_worker.len();
        
        for &ability in &sorted_worker {
            while job_index < n && ability >= jobs[job_index].0 {
                best_profit = std::cmp::max(best_profit, jobs[job_index].1);
                job_index += 1;
            }
            max_profit += best_profit;

        } 

        max_profit
    }
}
