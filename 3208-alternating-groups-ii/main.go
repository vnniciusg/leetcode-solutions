func numberOfAlternatingGroups(colors []int, k int) int {
    n := len(colors)
    ans, l := 0, 0

    for r := 1; r < n + k - 1; r++ {
        
        if colors[r % n] == colors[(r - 1) % n] {
            l = r
        }

        if r - l + 1 > k {
            l += 1
        }

        if r - l + 1 == k {
            ans++
        }

    }

    return ans
}
