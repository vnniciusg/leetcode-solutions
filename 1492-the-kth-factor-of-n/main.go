func kthFactor(n int, k int) int {
    
    var factorList []int
    for i := 1; i <= n; i++ {
        if n % i == 0 {
            factorList = append(factorList, i)
        }
    }

    if k - 1 >= len(factorList) { return - 1}

    return factorList[k - 1]
}
