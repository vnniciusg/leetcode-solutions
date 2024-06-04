func largestPalindromic(num string) string {
	count := make(map[rune]int)
	for _, digit := range num {
		count[digit]++
	}

	var digits []rune
	for digit := range count {
		digits = append(digits, digit)
	}
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] > digits[j]
	})

	firstHalf := []rune{}
	middle := rune(0)

	for _, digit := range digits {
		if count[digit]%2 != 0 {
			if middle == 0 || digit > middle {
				middle = digit
			}
		}
		firstHalf = append(firstHalf, []rune(strings.Repeat(string(digit), count[digit]/2))...)
	}

	firstHalfStr := strings.TrimLeft(string(firstHalf), "0")
	if firstHalfStr == "" {
		if middle != 0 {
			return string(middle)
		}
		return "0"
	}

	palindrome := firstHalfStr
	if middle != 0 {
		palindrome += string(middle)
	}
	palindrome += reverseString(firstHalfStr)
	return palindrome
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
