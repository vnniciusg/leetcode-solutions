func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    
    var n int = len(customers)

    var satisfiedCustomers int = 0
    for idx, costumer := range customers {
        if grumpy[idx] == 0 {
            satisfiedCustomers += costumer
        }
    }

    var additionalSatisfied int = 0
    for idx := range minutes {
        if grumpy[idx] == 1 {
            additionalSatisfied += customers[idx]
        }
    }

    var maxAdditionalSatisfied int = additionalSatisfied

    for idx := minutes; idx < n; idx++ {
        if grumpy[idx] == 1 {
            additionalSatisfied += customers[idx]
        }

        if grumpy[idx - minutes] == 1 {
            additionalSatisfied -= customers[idx - minutes]
        }

        maxAdditionalSatisfied = max(maxAdditionalSatisfied, additionalSatisfied)

    }

    return maxAdditionalSatisfied + satisfiedCustomers
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
