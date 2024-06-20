class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        
        max_ones = curr_ones = 0

        for num in nums:
            if num == 1:
                curr_ones += 1
                max_ones = max(max_ones, curr_ones)
            else:
                curr_ones = 0
        
        return max_ones
