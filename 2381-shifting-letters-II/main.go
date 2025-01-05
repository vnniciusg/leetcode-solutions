func shiftingLetters(s string, shifts [][]int) string {
    n := len(s)
    diff := make([]int, n+1)

    for _, shift := range shifts {
        start, end, direction := shift[0], shift[1], shift[2]
        if direction == 1 {
            diff[start]++
            diff[end+1]--
        }else {
            diff[start]--
            diff[end+1]++
        }
    }

    result := []rune(s)
    currShift := 0

    for i := 0; i < n; i++ {
        currShift = ((currShift + diff[i]) % 26 + 26) % 26
        curr := int(result[i] - 'a')
        newChar := (curr + currShift) % 26
        result[i] = rune(newChar + 'a')
    }

    return string(result)
}
