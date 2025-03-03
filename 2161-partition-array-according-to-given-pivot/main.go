func pivotArray(nums []int, pivot int) []int {
    
    var left, right, mid []int

    for _, num := range nums {
        if num < pivot {
            left = append(left, num)
        }else if num > pivot {
            right = append(right, num)
        }else {
            mid = append(mid, num)
        }
    }

    return append(left, append(mid, right...)...)
}
