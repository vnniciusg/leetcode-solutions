package main

func numJewelsInStones(jewels string, stones string) int {

    stonesMap := make(map[rune]int)
    for _, stone := range stones {
        stonesMap[stone]++
    }

    var count int
    for _, jewel := range jewels {
        count += stonesMap[jewel]
    }

    return count

}
