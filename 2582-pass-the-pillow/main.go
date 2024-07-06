func passThePillow(n int, time int) int {
    
    position, direction := 1, 1

    for range time {
        position += direction
        if position == n {
            direction = -1
        }else if position == 1 {
            direction = 1
        } 
    }

    return position
}
