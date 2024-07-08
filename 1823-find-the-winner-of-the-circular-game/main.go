func findTheWinner(n int, k int) int {
    friends := make([]int, n)
    for i := range n {
        friends[i] = i + 1
    }

    idx := 0

    for len(friends) > 1 {
        idx = (idx + k - 1) % len(friends)
        friends = append(friends[:idx], friends[idx+1:]...)
    }

    return friends[0]
}