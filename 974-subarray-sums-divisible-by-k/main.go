package main

func subarraysDivByK(nums []int, k int) int {
    
    count, prefixSum := 0, 0
    modCount := make(map[int]int)
    modCount[0] = 1

    for _, num := range nums {
        prefixSum += num
        mod := prefixSum % k

        if mod < 0 {
            mod += k
        }

        count += modCount[mod]
        modCount[mod]++
    }

    return count

}
