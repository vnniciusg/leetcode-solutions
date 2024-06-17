func wordBreak(s string, wordDict []string) []string {

    n := len(s)
    var sentences []string

    var backtrack func(start int, path []string)
    backtrack = func (start int, path []string){
        if start == n {
            sentences = append(sentences, strings.Join(path, " "))
            return
        }

        for end := start + 1; end < n + 1; end++{
            if isInWordDict(wordDict, s[start:end]){
                backtrack(end, append(path, s[start:end]))
            } 
        }
    }

    backtrack(0, []string{})
    return sentences
}

func isInWordDict(wordDict []string, word string) bool {
    for _, w := range wordDict {
        if w == word{
            return true
        }
    } 
    return false
}