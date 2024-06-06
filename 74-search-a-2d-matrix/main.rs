impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {

        if matrix.is_empty() || matrix[0].is_empty() {
            return false;
        }
        
        let m = matrix.len();
        let n = matrix[0].len();

        let mut left = 0;
        let mut right = (m * n) - 1;
        
        while left <= right {
            let mid = left + (right - left) / 2;
            let i = mid / n;
            let j = mid % n;

            if i >= m || j >= n {
                return false;
            }

            if target == matrix[i][j]{
                return true;
            }else if target > matrix[i][j] {
                left = mid + 1;
            }else {
                right = mid - 1;
            }
        } 

        false
    }
}
