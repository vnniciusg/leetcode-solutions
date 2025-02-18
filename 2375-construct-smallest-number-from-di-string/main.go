func smallestNumber(pattern string) string {
    n := len(pattern)
    var ans []string
    var numStack []int
    
    for idx := range n + 1 {

        numStack = append(numStack, idx + 1)

        if idx == n || pattern[idx] == 'I' {
            for len(numStack) > 0 {
                top := numStack[len(numStack) - 1]
                numStack = numStack[:len(numStack) - 1]
                ans = append(ans, strconv.Itoa(top))
            }
        }
    }

    return strings.Join(ans, "")
}
