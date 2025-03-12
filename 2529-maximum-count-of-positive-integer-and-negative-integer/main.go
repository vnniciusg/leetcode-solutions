func maximumCount(nums []int) int {
    return max(binarySearch(nums, 0), len(nums) - binarySearch(nums, 1))
}

func binarySearch(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    result := len(nums)

    for l <= r {

        mid := l + (r - l) / 2

        if nums[mid] < target {
            l = mid + 1
        }else {
            r = mid - 1
            result = mid
        }
    }

    return result
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
