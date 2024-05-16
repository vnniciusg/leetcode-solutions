package main

func getWinner(arr []int, k int) int {

	if k >= len(arr) {
		max := arr[0]
		for _, v := range arr {
			if v > max {
				max = v
			}
		}
		return max
	}

	winner := arr[0]
	winCount := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] > winner {
			winner = arr[i]
			winCount = 1
		} else {
			winCount += 1
		}

		if winCount == k {
			return winner
		}
	}

	return winner
}
