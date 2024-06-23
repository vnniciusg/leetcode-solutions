func diagonalPrime(nums [][]int) int {

    if len(nums) == 0 {
        return 0
    }

    var primeNumbers []int
    for i := range nums {
        primeNumbers = append(primeNumbers, nums[i][i])
        primeNumbers = append(primeNumbers, nums[i][len(nums) - i - 1])
    }

    return maxPrime(primeNumbers)
}

func maxPrime(primeNumbers []int) int {
    maxPrime := 0
    for _, num := range primeNumbers {
        if isPrime(num) && num > maxPrime {
            maxPrime = num
        }
    }

    return maxPrime
}

func isPrime(num int) bool {
    
    if num <= 1 {
        return false
    }

    for i := 2; i < int(math.Pow(float64(num), 0.5)) + 1; i++ {
        if num % i == 0 {
            return false
        }
    }

    return true
}
