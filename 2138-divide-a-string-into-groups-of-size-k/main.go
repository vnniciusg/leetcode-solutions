func divideString(s string, k int, fill byte) []string {
    
    var splitStringList []string

	for i := 0; i < len(s); i += k {
		end := i + k
		if end > len(s) {
			end = len(s)
		}
		substr := s[i:end]
		if len(substr) != k {
			substr += strings.Repeat(string(fill), k-len(substr))
		}
		splitStringList = append(splitStringList, substr)
	}

	return splitStringList
}
