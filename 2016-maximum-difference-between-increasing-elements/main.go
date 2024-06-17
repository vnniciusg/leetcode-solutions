func maximumDifference(nums []int) int {
    
    if nums == nil {
        return 0
    }

    maxDifference := - 1
    minValue := nums[0] 

    for _, num := range nums {
        if num > minValue {
            maxDifference = max(maxDifference, num - minValue)
        }else {
            minValue = num
        }
    }


    return maxDifference
}

func max(a, b int) int {
    if a < b {
        return b
    }

    return a
}
