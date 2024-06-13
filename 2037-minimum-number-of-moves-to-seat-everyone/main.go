func minMovesToSeat(seats []int, students []int) int {
    sort.Ints(seats)
    sort.Ints(students)

    var totalMoves int = 0
    m := len(seats)
    for i := range(m) {
        totalMoves += abs(seats[i] - students[i])
    }

    return totalMoves
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
