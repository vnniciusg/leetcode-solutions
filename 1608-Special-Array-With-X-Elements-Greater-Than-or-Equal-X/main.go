package main

func specialArray(nums []int) int {

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    n := len(nums)

    for i := 1; i < n + 1; i++{
        count := 0
        for _, num := range nums {
            if num >= i{
                count++
            }
        }
        if count == i {
            return i
        }
    }

    return -1
}
