package main

func stringMatching(words []string) []string {
    str := strings.Join(words, " ")
    var ans []string

    for _, word := range words {
        if strings.Count(str, word) > 1 {
            ans = append(ans, word)
        }
    }

    return ans
}
