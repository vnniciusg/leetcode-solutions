package main

func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0)
	left, right, top, bottom := 0, n-1, 0, m-1

	for left <= right && top <= bottom {

		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top <= bottom {
			for j := right; j >= left; j-- {
				res = append(res, matrix[bottom][j])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res
}
