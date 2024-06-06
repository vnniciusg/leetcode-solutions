package main


func searchInsert(nums []int, target int) int {

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    left, right := 0, len(nums) - 1
    for left <= right {
        mid := left + (right - left) / 2
        if nums[mid] == target {
            return mid
        }else if nums[mid] < target {
            left = mid + 1
        }else {
            right = mid -1
        }
    }
    return left
}

