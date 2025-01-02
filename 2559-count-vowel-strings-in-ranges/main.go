package main


func vowelStrings(words []string, queries [][]int) []int {

    var vowels = map[byte]struct{}{
        'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
        'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
    }

    prefixSum := make([]int, len(words))

    isVowel := func(c byte) bool {
        _, ok := vowels[c]
        return ok
    }

    if len(words[0]) > 0 && isVowel(words[0][0]) && isVowel(words[0][len(words[0])-1]) {
        prefixSum[0] = 1
    }

    for i := 1; i < len(words); i++ {
        prefixSum[i] = prefixSum[i-1]
        if len(words[i]) > 0 && isVowel(words[i][0]) && isVowel(words[i][len(words[i])-1]) {
            prefixSum[i]++
        }
    }

    result := make([]int, len(queries))
    for i, query := range queries {
        start, end := query[0], query[1]
        
        if start < 0 || end >= len(words) || start > end {
            continue
        }

        if start == 0 {
            result[i] = prefixSum[end]
        } else {
            result[i] = prefixSum[end] - prefixSum[start-1]
        }
    }

    return result

}
