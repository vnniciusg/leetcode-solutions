func maxSubArray(nums []int) int {
    
    curSum, maxSum := 0, int(math.Inf(-1))
    n := len(nums)

    for i := range(n){
        curSum += nums[i]
        maxSum = max(maxSum, curSum)
        if curSum < 0 {
            curSum = 0
        }
    }

    return maxSum
}

func max(x, y int) int {
    if x < y { return y }
    return x
}
