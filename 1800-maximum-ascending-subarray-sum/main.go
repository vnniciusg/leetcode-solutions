func maxAscendingSum(nums []int) int {
    _max := 0
    currSubSum := nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] <= nums[i - 1] {
            _max = max(_max, currSubSum)
            currSubSum = 0
        }

        currSubSum += nums[i]
    }

    return max(_max, currSubSum)
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
