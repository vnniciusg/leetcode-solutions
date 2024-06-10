func heightChecker(heights []int) int {
    
    n := len(heights)

    expected := make([]int, n)
    copy(expected, heights)

    sort.Slice(expected, func(i, j int) bool {
        return expected[i] < expected[j]
    })

    count := 0
    for i := range n {
        if heights[i] != expected[i] {
            count++
        }
    }

    return count
}
