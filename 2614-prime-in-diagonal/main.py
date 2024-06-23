class Solution:
    def diagonalPrime(self, nums: List[List[int]]) -> int:
        
        def is_prime(num: int) -> bool:
            
            if num <= 1:
                return False
            
            for i in range(2, int(num ** 0.5) + 1):
                if num % i == 0:
                    return False
            
            return True
        
        prime_numbers = [nums[i][i] for i in range(len(nums))]
        prime_numbers += [nums[i][len(nums) - i - 1] for i in range(len(nums))]

        prime_numbers = [num for num in prime_numbers if is_prime(num)]

        return max(prime_numbers) if prime_numbers else 0
