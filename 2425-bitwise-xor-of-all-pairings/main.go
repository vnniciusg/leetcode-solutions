func xorAllNums(nums1 []int, nums2 []int) int {
    xor1, xor2 := 0, 0
    len1, len2 := len(nums1), len(nums2)

    if len2 & 1 != 0 {
        for _, num := range nums1 {
            xor1 ^= num
        }
    }

    if len1 & 1 != 0 {
        for _, num := range nums2 {
            xor2 ^= num
        }
    }

    return xor1 ^ xor2
}
