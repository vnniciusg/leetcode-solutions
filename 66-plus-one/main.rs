impl Solution {
    pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
        

        let n = digits.len();
        let mut carry = 1;

        let mut ans = digits.clone();

        for i in (0..=n-1).rev() {
            ans[i] += carry;
            carry = ans[i] / 10;
            ans[i] %= 10;
        }

        if carry != 0 {
            ans.insert(0, carry);
        }

        ans
    }
}
