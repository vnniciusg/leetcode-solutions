func plusOne(digits []int) []int {
    
    n := len(digits)
    carry := 1
    for i := n - 1; i >= 0; i-- {
        digits[i] += carry
        carry = digits[i] / 10
        digits[i] %= 10
    }

    if carry != 0 {
        digits = append([]int{carry}, digits...)
    }

    return digits
}
