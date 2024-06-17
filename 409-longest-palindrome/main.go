package main

func longestPalindrome(s string) int {

    n := len(s)
    freq := make(map[byte]int)

    for i := 0; i < n; i++{
        freq[s[i]] += 1
    }

    var length int = 0
    var oddFound bool = false

    for _, value := range freq {
        if value & 1 == 0{
            length += value
        }else{
            length += value - 1
            oddFound = true
        }
    }

    if oddFound {
        length += 1
    }

    return length
}
