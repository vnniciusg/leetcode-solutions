func largestAltitude(gain []int) int {
    
    maxAltitude, currAltitude := 0, 0

    for _, altitude := range gain {
        currAltitude += altitude
        maxAltitude = max(maxAltitude, currAltitude)
    }

    return maxAltitude
}


func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
