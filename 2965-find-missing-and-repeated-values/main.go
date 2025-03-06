func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    count := make(map[int]int)

    for _, row := range grid {
        for _, num := range row {
            count[num]++
        }
    }

    repeated, missing := -1, -1
    for num := 0; num <= n * n; num++ {
        if count[num] == 2 {
            repeated = num
        }else if count[num] == 0 {
            missing = num
        }
    }

    return []int{repeated, missing}
}
