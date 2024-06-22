func numberOfSubarrays(nums []int, k int) int {
    
    hashMap := make(map[int]int)
    hashMap[0] = 1
    count, oddCount := 0, 0

    for _, num := range nums {
        if num & 1 != 0 {
            oddCount++
        }

        if value, exists := hashMap[oddCount - k]; exists {
            count += value
        }

        hashMap[oddCount]++
    }

    return count
}
