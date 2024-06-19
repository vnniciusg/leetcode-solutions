func minDays(bloomDay []int, m int, k int) int {

    if k * m > len(bloomDay) {
        return - 1
    }

    left, right := 1, 1
    for _, day := range bloomDay {
        if day > right {
            right = day
        }
    }   

    for left < right {
        mid := (left + right) / 2
        if canMakeBouquets(bloomDay, mid, m, k) {
            right = mid
        }else {
            left = mid + 1
        }
    }

    if canMakeBouquets(bloomDay, left, m, k) {
        return left
    }

    return -1

}

func canMakeBouquets(bloomDay []int, day, m, k int) bool {
    
    bouquets, flowers := 0, 0

    for _, bloom := range bloomDay {
        
        if bloom <= day {
            flowers++
            if flowers == k {
                bouquets++
                flowers = 0
            }
        }else {
            flowers = 0
        }

    }

    return bouquets >= m
}

