func pivotIndex(nums []int) int {
    
    leftSum, rightSum := 0, sum(nums)

    for idx, num := range nums {
        rightSum -= num

        if rightSum == leftSum {
            return idx
        }

        leftSum += num
    }

    return -1
}


func sum(nums []int) int {
    _sum := 0

    for _, num := range nums {
        _sum += num
    }

    return _sum
}
