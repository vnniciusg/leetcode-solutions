package main

func partition(s string) [][]string {
	res := make([][]string, 0)
	dfs(s, 0, []string{}, &res)
	return res
}

func dfs(s string, start int, path []string, res *[][]string) {
	if start == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
			dfs(s, i+1, path, res)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
