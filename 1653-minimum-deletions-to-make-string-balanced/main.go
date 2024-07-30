func minimumDeletions(s string) int {
    
    bCount, deletions := 0, 0
    
    for _, char := range s {
        if char == 'b'{ 
            bCount++
        }else {
            deletions = min(deletions + 1, bCount)
        }
    }

    return deletions
}

func min(a, b int) int  {
    if a > b {
        return b
    }
    return a
}
