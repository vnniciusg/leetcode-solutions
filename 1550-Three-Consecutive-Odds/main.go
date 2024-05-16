package main

func threeConsecutiveOdds(arr []int) bool {

	n := len(arr)

	count := 0
	for i := 0; i < n; i++ {
		if arr[i]&1 != 0 {
			count++
		} else {
			count = 0
		}

		if count == 3 {
			return true
		}
	}

	return false
}
