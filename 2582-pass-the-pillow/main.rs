impl Solution {
    pub fn pass_the_pillow(n: i32, time: i32) -> i32 {
        
        let (mut position, mut direction) = (1,  1);

        for _ in 0..time{
            position += direction;
            if position == n {
                direction = -1;
            }else if position == 1 {
                direction = 1;
            }
        }

        position
    }
}
