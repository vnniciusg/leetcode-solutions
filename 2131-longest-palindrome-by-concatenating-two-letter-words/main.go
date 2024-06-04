func longestPalindrome(words []string) int {
    freq := make(map[string]int)
    ans := 0
    f := 0

    for _, word := range words {
        freq[word]++
    }

    for k, v := range freq {
        x, y := string(k[0]), string(k[len(k)-1])
        if x == y {
            if v & 1 != 0 && f == 0 {
                f = 2
            }
            ans += 2 * (v / 2)
        } else if _, ok := freq[y+x]; ok {
            ans += min(v, freq[y+x])
        }
    }

    return 2*ans + f
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
