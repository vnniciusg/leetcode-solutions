func countTriplets(arr []int) int {
    
    n := len(arr)
    prefix := make([]int, n+1)

    for i := range n {
        prefix[i + 1] = prefix[i] ^ arr[i]
    }

    var count int = 0
    for j := range n {
        for k := j + 1; k < n; k++{
            if prefix[j] == prefix[k+1]{
                count += ( k - j)
            }
        }
    }

    return count
}
