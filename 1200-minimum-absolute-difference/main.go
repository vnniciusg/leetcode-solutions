func minimumAbsDifference(arr []int) [][]int {
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })

    n := len(arr) - 1
    _min := int(math.MaxInt64)
    for i := range arr[:n] {
        _min = min(_min, abs(arr[i] - arr[i+1]))
    }

    var ans [][]int
    for i := range arr[:n]{
        if abs(arr[i] - arr[i+1]) == _min {
            ans = append(ans, []int{arr[i], arr[i+1]})
        }
    }

    return ans

}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
