func romanToInt(s string) int {
    
    romanDict := map[rune]int {
            'I':1,
            'V':5,
            'X':10,
            'L':50,
            'C':100,
            'D':500,
            'M':1000,
    }

    total, prev := 0, 0

    for i := len(s) - 1; i >= 0; i-- {
        
        if value, ok := romanDict[rune(s[i])]; ok {

            if value >= prev {
                total += value
            }else {
                total -= value
            }

            prev = value
        }
    }
    
    return total
}
