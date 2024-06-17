package main

func maxScoreWords(words []string, letters []byte, score []int) int {

	var scoreMap = make(map[byte]int)
	for i, s := range score {
		scoreMap[byte(i+97)] = s
	}

	var letterMap = make(map[byte]int)
	for _, l := range letters {
		letterMap[l] += 1
	}

	var wordMap = make(map[string]map[byte]int)
	for _, word := range words {
		var m = make(map[byte]int)
		for _, l := range word {
			m[byte(l)] += 1
		}
		wordMap[word] = m
	}

	var maxScore int = 0

	var backtrack func(start int, score int)
	backtrack = func(start int, score int) {

		if score > maxScore {
			maxScore = score
		}

		for i := start; i < len(words); i++ {
			var canAdd bool = true
			var newScore int = score

			for k, v := range wordMap[words[i]] {
				if letterMap[k] < v {
					canAdd = false
					break
				} else {
					newScore += v * scoreMap[k]
				}
			}

			if canAdd {
				for k, v := range wordMap[words[i]] {
					letterMap[k] -= v
				}
				backtrack(i+1, newScore)
				for k, v := range wordMap[words[i]] {
					letterMap[k] += v
				}
			}
		}
	}

	backtrack(0, 0)
	return maxScore
}
