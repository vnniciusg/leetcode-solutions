func reverseVowels(s string) string {
    
    vowels := "aeiouAEIOU"
	vowelSet := make(map[rune]bool)
	for _, v := range vowels {
		vowelSet[v] = true
	}

	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		if vowelSet[runes[left]] && vowelSet[runes[right]] {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
		if !vowelSet[runes[left]] {
			left++
		}
		if !vowelSet[runes[right]] {
			right--
		}
	}

	return string(runes)
}
