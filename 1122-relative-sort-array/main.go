package main

func relativeSortArray(arr1 []int, arr2 []int) []int {
    
    orderMap := make(map[int]int)
    for i, num := range arr2 {
        orderMap[num] = i
    }

    countMap := make(map[int]int)
    for _, num := range arr1 {
        countMap[num]++
    }

    var ans []int
    for _, num := range arr2 {
        if count, exists := countMap[num]; exists {
			for range count {
				ans = append(ans, num)
			}
			delete(countMap, num)
		}
    }

    var remaining []int
    for num, count := range countMap {
        for range count {
            remaining = append(remaining, num)
        }
    }

    sort.Ints(remaining)
    ans = append(ans, remaining...)

    return ans
}
