func minimumLength(s string) int {

    freq := make(map[rune]int, len(s))
    for _, chr := range s {
        freq[chr] += 1
    }

    deleteCount := 0
    for _, frequency := range freq {
        if frequency & 1 != 0 {
            deleteCount += frequency - 1
        }else {
            deleteCount += frequency - 2
        }
    }

    return len(s) - deleteCount

}
