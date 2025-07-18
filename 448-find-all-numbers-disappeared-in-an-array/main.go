func findDisappearedNumbers(nums []int) []int {
    
    for _, num := range nums {
        idx := abs(num) - 1
        if nums[idx] > 0 {
            nums[idx] = -nums[idx]
        }
    }

    var missing []int
    for idx := range len(nums) {
        if nums[idx] > 0 {
            missing = append(missing, idx + 1)
        }
    }

    return missing
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}
