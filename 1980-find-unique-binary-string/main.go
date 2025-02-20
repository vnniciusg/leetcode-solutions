func findDifferentBinaryString(nums []string) string {
    var ans []string
    n := len(nums)

    for i := range n {
        curr := string(nums[i][i])
        newLetter := "0"
        if curr == "0" {
            newLetter = "1"
        }

        ans = append(ans, newLetter)
    }

    return strings.Join(ans, "")
}
