func countPrefixSuffixPairs(words []string) int {

    count := 0
    for i := range words {
        for j := i + 1; j < len(words); j++ {
            if strings.HasPrefix(words[j], words[i]) && strings.HasSuffix(words[j], words[i]) {
                count++
            }
        }
    }

    return count
}
