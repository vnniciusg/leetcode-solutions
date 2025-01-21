func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i := 0; i < len(nums2); i++ {
        nums1[i + m] = nums2[i]
    }

    sort.Slice(nums1, func(i, j int) bool {
        return nums1[i] < nums1[j]
    })
}
