package main

func uniqueOccurrences(arr []int) bool {

    freq := make(map[int]int)

    for _, num := range arr {
        freq[num] += 1
    }

    occurrence := make(map[int]bool)

    for _, count := range freq {
        if _, ok := occurrence[count]; ok {
            return false
        }

        occurrence[count] = true
    }

    return true
}


