func minKBitFlips(nums []int, k int) int {
    
    n := len(nums)
    flips, flipsEffect := 0, 0
    isFlipped := make([]int, n)

    for idx := range n {
        
        if idx >= k {
            flipsEffect ^= isFlipped[idx - k]
        }

        if flipsEffect == nums[idx] {

            if idx + k > n {
                return -1
            }

            flips += 1
            flipsEffect ^= 1
            isFlipped[idx] = 1
        }
    }

    return flips
}
