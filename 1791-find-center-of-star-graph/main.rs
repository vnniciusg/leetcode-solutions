impl Solution {
    pub fn find_center(edges: Vec<Vec<i32>>) -> i32 {
        let (a, b) = (edges[0][0], edges[0][1]);

        if edges[1].contains(&a) {
            return a;
        }

        b
    }
}
