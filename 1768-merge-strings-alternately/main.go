func mergeAlternately(word1 string, word2 string) string {
    
    l1, l2 := len(word1), len(word2)
    minLength := min(l1, l2)

    var ans string = ""
    for i := range minLength {
        ans += string(word1[i]) + string(word2[i])
    }

    if l1 > l2 {
        ans += word1[minLength:]
    }else {
        ans += word2[minLength:]
    }

    return ans
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}
