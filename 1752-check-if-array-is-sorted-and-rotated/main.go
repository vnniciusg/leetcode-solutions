func check(nums []int) bool {
    n := len(nums)

    sorted_nums := make([]int, n)
    copy(sorted_nums, nums)
    sort.Ints(sorted_nums)

    for i := range nums {
        isMatch := true
        
        for j := range nums {
            if sorted_nums[j] != nums[(i + j) % n] {
                isMatch = false
                break
            }
        }

        if isMatch {
            return true
        }
    }

    return false
}
