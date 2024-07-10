func minOperations(logs []string) int {
    
    var depth int = 0

    for _, log := range logs {
        switch log {
        case "../":
            if depth > 0 {
                depth -= 1
            }
        case "./":
            continue
        default:
            depth += 1
        }
    }

    return depth
}
