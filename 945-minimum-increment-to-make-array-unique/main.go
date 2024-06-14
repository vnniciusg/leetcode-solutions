func minIncrementForUnique(nums []int) int {
    
    sort.Ints(nums)

    var increments int = 0
    for i := 1; i < len(nums); i++ {
        if nums[i] <= nums[i - 1] {
            needed := nums[i - 1] + 1  - nums[i]
            nums[i] += needed
            increments += needed
        }
    }

    return increments
}
