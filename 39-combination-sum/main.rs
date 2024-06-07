impl Solution {
    pub fn combination_sum(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        
        let n = candidates.len();

        let mut ans: Vec<Vec<i32>> = Vec::new();

        let mut sorted_candidates = candidates.clone();
        sorted_candidates.sort();

        dfs(&sorted_candidates, target, 0, &mut Vec::new(), &mut ans);

        ans
    }
}


fn dfs(
        candidates: &[i32],
        target: i32,
        index: usize,
        path: &mut Vec<i32>,
        result: &mut Vec<Vec<i32>>,
    ) {
        if target < 0 {
            return;
        }

        if target == 0 {
            result.push(path.clone());
            return;
        }

        for i in index..candidates.len() {
            path.push(candidates[i]);
            dfs(candidates, target - candidates[i], i, path, result);
            path.pop();
        }
    }
