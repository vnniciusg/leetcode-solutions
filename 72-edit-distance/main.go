func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := buildDistanceMatrix(m, n)

    for idx := 1; idx < m + 1; idx++ {
        for jdx := 1; jdx < n + 1; jdx++ {
            cost := 1
            if word1[idx - 1] == word2[jdx - 1] {
                cost = 0
            }
            dp[idx][jdx] = min(
                dp[idx - 1][jdx] + 1,
                dp[idx][jdx - 1] + 1,
                dp[idx - 1][jdx - 1] + cost,
            )
        }
    }

    return dp[m][n]
}

func buildDistanceMatrix(m, n int) [][]int {
    dp := make([][]int, m + 1)

    for idx := range dp {
        dp[idx] = make([]int, n + 1)
    }

    for idx := 0; idx <= m; idx++ {
        dp[idx][0] = idx
    }

    for jdx := 0; jdx <= n; jdx++ {
        dp[0][jdx] = jdx
    }

    return dp
}

func min(nums ...int) int {
    _min := math.MaxInt
    for _, num := range nums {
        if num < _min {
            _min = num
        }
    }
    return _min
}
