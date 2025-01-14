func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    result := make([]int, n)
    seenA := make(map[int]bool)
    seenB := make(map[int]bool)

    for i := 0; i < n; i++ {

        seenA[A[i]] = true
        seenB[B[i]] = true

        common := 0
        for num := range seenA {
            if seenB[num] { 
                common++
            }
        }

        result[i] = common
    }

    return result
}
