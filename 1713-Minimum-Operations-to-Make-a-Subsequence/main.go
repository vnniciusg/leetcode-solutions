package main

func minOperations(target []int, arr []int) int {

    targetIndex := make(map[int]int)
    for i, num := range target {
        targetIndex[num] = i
    }

    var targetIndices []int
    for _, num := range arr {
        if index, exists := targetIndex[num]; exists {
            targetIndices = append(targetIndices, index)
        }
    }

    lisLen := lenOfLis(targetIndices)
    return len(target) - lisLen
}

func lenOfLis(nums []int) int {
    lis := []int{}
    for _, num := range nums {
        pos := sort.SearchInts(lis, num)
        if pos < len(lis) {
            lis[pos] = num
        }else{
            lis = append(lis, num)
        }
    }

    return len(lis)
}
