func missingNumber(nums []int) int {
    result := 0

    for i := 1; i <= len(nums); i++ {
        result ^= i
    } 

    for _, value := range nums {
        result ^= value
    }

    return result
}
