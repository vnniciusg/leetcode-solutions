func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    
    n := len(difficulty)
    var jobs [][2]int
    for i := range n {
        jobs = append(jobs, [2]int{difficulty[i], profit[i]})
    }

    sort.Slice(jobs, func(i, j int) bool {
        return jobs[i][0] < jobs[j][0]
    })
    
    sort.Ints(worker)

    maxProfit, bestProfit := 0, 0
    jobIndex := 0
    n = len(worker)

    for _, ability := range worker {
        for jobIndex < n && ability >= jobs[jobIndex][0] {
            bestProfit = max(bestProfit, jobs[jobIndex][1])
            jobIndex++
        }

        maxProfit += bestProfit
    }

    return maxProfit
}

func max(i, j int) int {
    if i < j {
        return j
    }

    return i
}
