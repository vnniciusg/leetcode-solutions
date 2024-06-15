func kidsWithCandies(candies []int, extraCandies int) []bool {
    
    var ans []bool
    max := max(candies)
    for _, kid := range candies {
        ans = append(ans, (kid + extraCandies >= max))
    }

    return ans
}


func max(nums []int) int {
    maxValue := 0
    for _, num := range nums {
        if num > maxValue {
            maxValue = num
        }
    }
    return maxValue
}
