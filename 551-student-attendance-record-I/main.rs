impl Solution {
    pub fn check_record(s: String) -> bool {
        let absences = s.chars().filter(|&c| c == 'A').count();
        if absences >= 2 {
            return false;
        }

        if s.contains("LLL") {
            return false;
        }

        true
    }
}
