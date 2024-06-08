func construct2DArray(original []int, m int, n int) [][]int {
    
    if len(original) != m * n {
        return [][]int{}
    }

    var ans [][]int
    for i := range(m) {
        row := original[i * n:(i + 1) * n]
        ans = append(ans, row)
    }

    return ans
}
