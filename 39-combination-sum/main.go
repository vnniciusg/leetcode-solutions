func combinationSum(candidates []int, target int) [][]int {

    n := len(candidates)

    var dfs func(target int, index int, path []int, ans *[][]int) 
    dfs = func(target int, index int, path []int, ans *[][]int) {

        if target < 0 {
            return 
        }

        if target == 0 {
            *ans = append(*ans, append([]int{}, path...)) 
            return
        }

        for i := index; i < n; i++ {
            dfs(target-candidates[i], i, append(path, candidates[i]) ,ans)
        }
    }

    sort.Slice(candidates, func(i, j int) bool {
        return candidates[i] < candidates[j]
    })

    
    ans := [][]int{}
    dfs(target, 0, []int{}, &ans)
    return ans
}
