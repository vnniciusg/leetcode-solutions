func longestConsecutive(nums []int) int {
    
    numsSet := make(map[int]bool)
    for _, num := range nums {
        numsSet[num] = true
    }

    var longest int = 0
    for num := range numsSet {
            var length int = 1
            for numsSet[num + length] {
                length++
            }
            longest = max(longest, length)            
        }
    }

    return longest
}


func max(a, b int) int {
    if a < b {
        return b
    }
    

    return a
}
