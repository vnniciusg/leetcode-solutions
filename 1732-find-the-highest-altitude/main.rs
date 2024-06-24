impl Solution {
    pub fn largest_altitude(gain: Vec<i32>) -> i32 {
        
        let (mut max_altitude, mut curr_altitude) = (0, 0);

        for altitude in gain.iter() {
            curr_altitude += altitude;
            max_altitude = std::cmp::max(max_altitude, curr_altitude);
        }

        max_altitude
    }
}
