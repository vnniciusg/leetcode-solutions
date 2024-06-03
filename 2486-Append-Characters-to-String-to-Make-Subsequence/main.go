package main

func appendCharacters(s string, t string) int {
    
    n , m := len(s), len(t)
    
    i , j := 0, 0
    for i < n && j < m {
        if s[i] == t[j] {
            j += 1
        }
        i += 1
    }

    return m - j

}
