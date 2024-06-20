func findMaxConsecutiveOnes(nums []int) int {

    var max_sum int = 0
    var curr_sum int = 0 

    for _, num := range nums {
        if num == 1 {
            curr_sum++
            max_sum = max(max_sum, curr_sum)
        }else {
            curr_sum = 0
        }
    }
    return max_sum

}

func max(a,b int) int {
    if a < b {
        return b
    }
    return a
}
