func singleNumber(nums []int) []int {
    
    elements := make(map[int]int)
    for _, num := range nums {
        elements[num]++
    }

    var res []int
    for k, v := range elements {
        if v == 1 {
            res = append(res, k)
        }
    }

    return res
}
