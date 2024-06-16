func maxOperations(nums []int, k int) int {
    
    countMap := make(map[int]int)
    countPairs := 0

    for _, num := range nums {
        complement := k - num
        if countMap[complement] > 0 {
            countPairs++
            countMap[complement]--
        }else {
            countMap[num]++
        }
    }

    return countPairs
}
