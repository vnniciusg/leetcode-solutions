func scoreOfString(s string) int {
    
    n := len(s)
    var score int = 0

    for i := 0; i < n -1; i++{
        score += abs(int(s[i]) - int(s[i+1])) 
    }

    return score
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
