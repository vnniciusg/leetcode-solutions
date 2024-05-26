package main

import "strings"

func checkRecord(s string) bool {
	absences := strings.Count(s, "A")
	if absences >= 2 {
		return false
	}

	if strings.Contains(s, "LLL") {
		return false
	}

	return true
}
