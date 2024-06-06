func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])

    left, right := 0, m * n -1
    for left <= right {
        mid := left + (right - left) / 2
        i, j := mid / n, mid % n
        if target == matrix[i][j]{
            return true
        }else if target  > matrix[i][j]{
            left = mid + 1
        }else {
            right = mid -1
        }
    }

    return false
}
