package main

func smallestDistancePair(nums []int, k int) int {

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    var countPairs func(mid int) int
    countPairs = func(mid int) int  {
        n := len(nums)
        count, left := 0, 0

        for right := range(n) {
            for nums[right] - nums[left] > mid {
                left++
            }
            count += right - left
        }

        return count
    }

    left, right := 0, nums[len(nums)-1]-nums[0]
    for left < right {
        mid := left + (right - left) / 2
        if countPairs(mid) < k {
            left = mid + 1
        }else {
            right = mid
        }
    }

    return left
}

