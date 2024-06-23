func judgeCircle(moves string) bool {
    
    position := [2]int{0, 0}

    for _, move := range moves {
        if move == 'U' {
            position[0]++
        }else if move == 'D' {
            position[0]--
        }else if move == 'L'{
            position[1]--
        }else {
            position[1]++
        }
    }

    return position == [2]int{0, 0}
}
