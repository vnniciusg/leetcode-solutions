func singleNumber(nums []int) int {
    
    var xor int = 0
    for _, num := range nums {
        xor ^= num
    }

    return xor
}
