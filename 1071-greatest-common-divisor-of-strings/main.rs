impl Solution {
    pub fn gcd_of_strings(str1: String, str2: String) -> String {
        if str1.to_owned() + &str2 != str2.to_owned() + &str1 {
            return String::from("");
        }

        let gcd_length = gcd(str1.len(), str2.len());
        str1[..gcd_length].to_string()
    }
}

fn gcd(mut a: usize, mut b: usize) -> usize {

    while b != 0 {
        let temp = b;
        b = a % b;
        a = temp;
    }

    a
}
