package main

func areAlmostEqual(s1 string, s2 string) bool {

	if s1 == s2 {
		return true
	}

	diffCount := 0
	var diffIndeces []int

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffCount += 1
			diffIndeces = append(diffIndeces, i)

			if diffCount > 2 {
				return false
			}
		}
	}

	return diffCount == 2 && s1[diffIndeces[0]] == s2[diffIndeces[1]] && s1[diffIndeces[1]] == s2[diffIndeces[0]]

}
