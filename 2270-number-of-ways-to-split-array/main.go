func waysToSplitArray(nums []int) int {
    left := 0
    right := sum(nums)

    count := 0
    for idx := 0; idx < len(nums) - 1; idx++ {
        left += nums[idx]
        right -= nums[idx]

        if left >= right {
            count++
        }
    }

    return count
}


func sum(nums []int) int {  
    _sum := 0

    for _, num := range nums{
        _sum += num
    }
    return _sum
}
