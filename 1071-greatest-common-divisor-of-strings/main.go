func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }

    gcdLength := gcd(len(str1), len(str2))
	return str1[:gcdLength]
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
