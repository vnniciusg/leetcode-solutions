class Solution:
    def distinctPrimeFactors(self, nums: List[int]) -> int:
        
        def prime_factors(n: int, i: int = 2, factors=None):
            
            if factors is None:
                factors = set()
            
            if n <= 1:
                return factors

            if n % i == 0:
                factors.add(i)
                return prime_factors(n // i, i, factors)
            
            return prime_factors(n, i+1, factors)
        

        distinct_primes = set()
        for num in nums:
            distinct_primes.update(prime_factors(num))
        

        return len(distinct_primes)
