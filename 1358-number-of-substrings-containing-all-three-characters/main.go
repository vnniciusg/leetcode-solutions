func numberOfSubstrings(s string) int {
    
    l, ans := 0, 0
    count := make(map[byte]int)

    n := len(s)
    for r := 0; r < n; r++ {
        count[s[r]]++

        for len(count) == 3 {
            ans += (n - r)
            count[s[l]]--
            if count[s[l]] == 0 {
                delete(count, s[l])
            }
            l++
        }
    }

    return ans
}
