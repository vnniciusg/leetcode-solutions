impl Solution {
    pub fn num_steps(s: String) -> i32 {        
        let mut s = s.chars().collect::<Vec<char>>();
        let mut steps = 0;
        while s.len() > 1 {
            if s[s.len() - 1] == '0' {
                s.pop();
            } else {
                let mut carry = true;
                for i in (0..s.len()).rev() {
                    if s[i] == '1' {
                        if carry {
                            s[i] = '0';
                        } else {
                            s[i] = '1';
                            break;
                        }
                    } else {
                        if carry {
                            s[i] = '1';
                            carry = false;
                        }
                    }
                }
                if carry {
                    s.insert(0, '1');
                }
            }
            steps += 1;
        }
        steps
    }
}