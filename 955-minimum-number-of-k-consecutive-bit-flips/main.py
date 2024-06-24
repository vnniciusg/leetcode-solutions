class Solution:
    def minKBitFlips(self, nums: List[int], k: int) -> int:
        
        n = len(nums)
        flips, flips_effect = 0, 0
        is_flipped = [0] * n

        for i in range(n):
            if i >= k:
                flips_effect ^= is_flipped[i - k]
            
            if flips_effect == nums[i]:

                if i + k > n:
                    return -1
                
                flips += 1
                flips_effect ^= 1
                is_flipped[i] = 1


        return flips
