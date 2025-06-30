func findLHS(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    maxLenght := 0
    j := 0

    for i := range nums {
        for nums[i] - nums[j] > 1 {
            j++
        }

        if nums[i] - nums[j] == 1 {
            maxLenght = max(maxLenght, i - j + 1)
        }
    }

    return maxLenght
}


func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}
