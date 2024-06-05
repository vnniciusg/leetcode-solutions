package main

func commonChars(words []string) []string {

	if len(words) == 0 {
		return []string{}
	}

	if len(words) == 1 {
		return []string{words[0]}
	}

	commonFreq := counter(words[0])

	for _, word := range words[1:] {
		wordFreq := counter(word)

		for char := range commonFreq {
			if count, ok := wordFreq[char]; ok {
				commonFreq[char] = min(commonFreq[char], count)
			} else {
				commonFreq[char] = 0
			}
		}
	}

	ans := []string{}
	for char, freq := range commonFreq {
		for range freq {
			ans = append(ans, string(char))
		}
	}

	return ans

}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func counter(word string) map[rune]int {

	freq := make(map[rune]int)

	for _, char := range word {
		freq[char]++
	}

	return freq
}
