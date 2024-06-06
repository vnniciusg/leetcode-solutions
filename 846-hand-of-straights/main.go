package main

func isNStraightHand(hand []int, groupSize int) bool {

    if len(hand) % groupSize != 0 {
        return false
    }

    sort.Slice(hand, func(i, j int) bool {
        return hand[i] < hand[j]
    })

    count := make(map[int]int)
    for _, card := range hand {
        count[card]++
    }

    for _, card := range hand {
        if count[card] > 0 {
            for i := range(groupSize) {
                if count[card + i] > 0 {
                    count[card + i] -= 1
                }else {
                    return false
                }
            }
        }
    }

    return true

}
