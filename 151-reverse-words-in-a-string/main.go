func reverseWords(s string) string {
    
    l := strings.Fields(s)
    n := len(l)

    for i := range (n / 2) {
        l[i] , l[n - i - 1] = l[n - i - 1], l[i]
    }

    return strings.Join(l, " ")
}
