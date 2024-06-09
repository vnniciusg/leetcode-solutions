package main

func subarraySum(nums []int, k int) int {
    
    ans, prefixSum := 0, 0
    count := make(map[int]int)
    count[0] = 1

    for _, num := range nums {
        prefixSum += num
        if value, ok := count[prefixSum - k]; ok {
            ans += value
        }
        
        count[prefixSum]++
    }

    return ans
}

