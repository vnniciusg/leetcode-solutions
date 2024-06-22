func numIdenticalPairs(nums []int) int {

    countValues := make(map[int]int)

    var countGoodPairs int = 0
    for _, num := range nums {
        countGoodPairs += countValues[num]
        countValues[num]++
    }
    
    return countGoodPairs
}
