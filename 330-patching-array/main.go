func minPatches(nums []int, n int) int {
    
    sum, cnt, i := 0, 0, 0

    for sum < n  {
        if i < len(nums) && nums[i] <= sum + 1 {
            sum += nums[i]
            i++
        }else {
            cnt++
            sum += sum + 1
        }
    }

    return cnt
}
