func increasingTriplet(nums []int) bool {
    
    if len(nums) < 3 {
        return false
    }

    first, second := math.MaxInt64, math.MaxInt64

    for _, num := range nums {
        if num <= first {
            first = num
        }else if num <= second {
            second = num
        }else {
            return true
        }
    }

    return false
}
