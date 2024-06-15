func twoSum(nums []int, target int) []int {
    
    _map := make(map[int]int)
    for idx, num := range nums {
        if value, ok := _map[num]; ok {
            return []int{idx, value}
        } 
        
        _map[target-num] = idx
    }

    return []int{}
}
