func countBadPairs(nums []int) int64 {
    
    n := len(nums)
    var totalPairs int64 = int64(n * (n-1) / 2)
    var goodPairs int64 = 0

    diffCount := make(map[int]int)

    for i, num := range nums {
        diff := num - i
        if value, ok := diffCount[diff]; ok {
            goodPairs += int64(value)
        }
        diffCount[diff]++
    }

    return totalPairs - goodPairs
}
