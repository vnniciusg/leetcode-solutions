impl Solution {
    pub fn min_moves_to_seat(seats: Vec<i32>, students: Vec<i32>) -> i32 {
        
        let mut sorted_seats = seats.clone();
        sorted_seats.sort();

        let mut sorted_students = students.clone();
        sorted_students.sort();

        let mut total_moves: i32 = 0;
        let m: usize = seats.len();

        for i in 0..m {
            total_moves += (sorted_seats[i] - sorted_students[i]).abs();
        }

        total_moves
    }
}
