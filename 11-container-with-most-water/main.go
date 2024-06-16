func maxArea(height []int) int {
    
    maxAmount := 0
    i, j := 0, len(height) - 1

    for i < j {
        currAmount := min(height[i], height[j]) * (j - i)

        if currAmount > maxAmount {
            maxAmount = currAmount
        }else if height[i] < height[j] {
            i++
        }else {
            j--
        }
    }

    return maxAmount
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
