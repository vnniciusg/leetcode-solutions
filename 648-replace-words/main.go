package main

func replaceWords(dictionary []string, sentence string) string {

    rootSet := make(map[string]bool)
    for _, word := range dictionary {
        rootSet[word] = true
    }

    sentenceList := strings.Split(sentence, " ")

    var ans []string
    for _, word := range sentenceList {
        ans = append(ans, findRoot(word, rootSet))
    }

    return strings.Join(ans, " ")

}

func findRoot(word string, rootSet map[string]bool) string {

    for i := 1; i < len(word) + 1; i++ {
        prefix := word[:i]
        if _, ok := rootSet[prefix]; ok {
            return prefix
        }
    }
    return word
}
