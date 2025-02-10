func clearDigits(s string) string {
	var ans []rune

	for _, char := range s {
		if unicode.IsDigit(char) && len(ans) > 0 {
			ans = ans[:len(ans)-1]
		} else {
			ans = append(ans, char)
		}
	}

	return string(ans)
}
