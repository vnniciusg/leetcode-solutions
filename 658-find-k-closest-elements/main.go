func findClosestElements(arr []int, k int, x int) []int {
    
    left, right := 0, len(arr) - 1

    for right - left >= k {
        if abs(arr[left] - x) > abs(arr[right] - x) {
            left++
        }else {
            right--
        }
    }

    return arr[left: right + 1]
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
