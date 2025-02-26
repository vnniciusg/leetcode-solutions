func maxAbsoluteSum(nums []int) int {
    var minSum, maxSum, minEnding, maxEnding int

    for _, num := range nums {
        minEnding = min(minEnding + num, num)
        maxEnding = max(maxEnding + num, num)
        
        minSum = min(minSum, minEnding)
        maxSum = max(maxSum, maxEnding)
    }

    return max(maxSum, abs(minSum))
}


func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func min(a, b int) int  {
    if a > b {
        return b
    }
    return a
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
