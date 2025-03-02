func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    
    _dict :=  make(map[int]int)

    for _, nums := range nums1 {
        _dict[nums[0]] = nums[1] 
    }

    for _, nums := range nums2 {
        _dict[nums[0]] += nums[1]
    }

    var ans [][]int 
    for key, value := range _dict {
        ans = append(ans, []int{key, value})
    }

    sort.Slice(ans, func(i, j int) bool {
        return ans[i][0] < ans[j][0]
    })

    return ans
}
