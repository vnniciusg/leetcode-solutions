class Solution:
    def maxAbsoluteSum(self, nums: List[int]) -> int:
        
        max_sum = min_sum = 0
        max_ending = min_ending = 0

        for num in nums:
            
            max_ending = max(max_ending + num, num)      
            min_ending = min(min_ending + num, num)
                  
            max_sum = max(max_sum, max_ending)
            min_sum = min(min_sum, min_ending)

        return max(max_sum, abs(min_sum))
