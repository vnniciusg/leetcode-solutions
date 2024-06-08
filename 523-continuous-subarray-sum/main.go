func checkSubarraySum(nums []int, k int) bool {
    
    if len(nums) < 2 {
        return false
    }

    remainderMap := make(map[int]int)
    remainderMap[0] = -1

    var cumulativeSum int = 0
    for i, num := range nums {
        cumulativeSum += num
        remainder := cumulativeSum % k

        if value, ok := remainderMap[remainder]; ok {
            if i - value > 1 {
                return true
            } 
        }else {
            remainderMap[remainder] = i
        }
    }

    return false

}

