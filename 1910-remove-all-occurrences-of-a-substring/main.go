func removeOccurrences(s string, part string) string {
    return helper(s, part)
}

func helper(substr string, part string) string {
    if !strings.Contains(substr, part) {
        return substr
    }

    substr = strings.Replace(substr, part, "", 1)
    substr = helper(substr, part)
    return substr
}
