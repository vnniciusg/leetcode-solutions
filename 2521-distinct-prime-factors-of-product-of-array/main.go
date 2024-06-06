package main

func distinctPrimeFactors(nums []int) int {
    distinctPrimes := make(map[int]bool)
    for _, num := range nums {
        factors := primeFactors(num)
        for factor := range factors {
            distinctPrimes[factor] = true
        }
    }
    return len(distinctPrimes)
}

func primeFactors(n int) map[int]bool {

    factors := make(map[int]bool)
    i := 2
    for n > 1 {
        if n%i == 0 {
            factors[i] = true
            n /= i
        }else{
            i++
        }
    }

    return factors
}
