func tupleSameProduct(nums []int) int {
    
    n := len(nums)
    freqMap := make(map[int]int)

    for i := 0; i < n - 1; i++ {
        for j := i + 1; j < n; j++ {
            freqMap[nums[i] * nums[j]]++
        }
    }

    count := 0
    for _, freq := range freqMap {
        count += 8 * (freq * (freq - 1)) / 2
    }

    return count
}
