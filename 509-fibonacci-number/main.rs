impl Solution {
    pub fn fib(n: i32) -> i32 {
        if n <= 1 {
            return n;
        }

        Self::fib(n-1) + Self::fib(n-2)
    }
}
