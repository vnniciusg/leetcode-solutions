class Solution:
    def minIncrementForUnique(self, nums: List[int]) -> int:
        
        nums.sort()
    
        increments = 0
        
        for i in range(1, len(nums)):
            if nums[i] <= nums[i - 1]:
                needed = nums[i - 1] + 1 - nums[i]
                nums[i] += needed
                increments += needed
        
        return increments
