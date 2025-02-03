func longestMonotonicSubarray(nums []int) int {
    incLength, decLength, maxLength := 1, 1, 1

    for i := 0; i < len(nums) - 1; i++ {
        if nums[i + 1] > nums[i] {
            incLength++
            decLength = 1
        }else if nums[i + 1] < nums[i] {
            decLength++
            incLength = 1
        }else {
            incLength = 1
            decLength = 1
        }

        maxLength = max(maxLength, max(incLength, decLength))
    }

    return maxLength
}

func max(a, b int) int {
    if a < b {
        return b
    }

    return a
}
