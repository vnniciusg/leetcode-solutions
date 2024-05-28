package main

func equalSubstring(s string, t string, maxCost int) int {

	var maxLen int = 0
	var cost int = 0
	var left int = 0
	var right int = 0

	for right < len(s) {
		cost += abs(int(s[right]) - int(t[right]))
		for cost > maxCost {
			cost -= abs(int(s[left]) - int(t[left]))
			left++
		}
		maxLen = max(maxLen, right-left+1)
		right++
	}
	return maxLen

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
