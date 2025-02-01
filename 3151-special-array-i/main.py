class Solution:
    def isArraySpecial(self, nums: List[int]) -> bool:
        
        n = len(nums)
        for i in range(1, n):
            par_prev, par_curr = nums[i - 1] & 1, nums[i] & 1
            if par_prev == par_curr: 
                return False

        return True
