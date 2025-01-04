package main

import "strings"

func countPalindromicSubsequence(s string) int {
    letters := make(map[rune]bool, len(s))
    for _, chr := range s {
        letters[chr] = true
    }

    ans := 0

    for letter := range letters {
        i := strings.IndexRune(s, letter)
        j := - 1
        for k := len(s) - 1; k >= 0; k-- {
            if rune(s[k]) == letter {
                j = k
                break
            }
        }

        if i == j || j == -1{
            continue
        }

        between := make(map[rune]bool)
        for _, chr := range s[i+1: j] {
            between[chr] = true 
        }

        ans += len(between)
    }


    return ans
}
